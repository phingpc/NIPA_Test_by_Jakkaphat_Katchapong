// KanbanColumn.jsx

import React, { useState, useRef, useEffect } from 'react';
import { useDrop } from 'react-dnd';
import KanbanCard from './KanbanCard.jsx';
import axios from 'axios';
import '../styles/KanbanColumn.css';

const KanbanColumn = ({ status, cards, onDrop, onCardUpdate }) => {
    const [isAdding, setIsAdding] = useState(false);
    const [newCardData, setNewCardData] = useState({
        title: '',
        description: '',
        contact: '',
    });
    const [error, setError] = useState(null);

    const formRef = useRef(null);
    const cardRefs = useRef([]);

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (formRef.current && !formRef.current.contains(event.target)) {
                handleAddCard();
                setIsAdding(false);
            }
        };
    
        if (isAdding) {
            document.addEventListener('mousedown', handleClickOutside);
        } else {
            document.removeEventListener('mousedown', handleClickOutside);
        }
    
        return () => document.removeEventListener('mousedown', handleClickOutside);
    }, [isAdding, newCardData]);

    const [, drop] = useDrop({
        accept: 'CARD',
        drop: (item) => onDrop(item.id, status),
    });

    const handleInputChange = (e) => {
        setNewCardData({ ...newCardData, [e.target.name]: e.target.value });
        setError(null);
    };

    const handleKeyDown = (e) => {
        if (e.key === "Enter") {
            handleAddCard();
        } else if (e.key === "Escape") {
            setIsAdding(false);
        }
    };

    const handleAddCard = async () => {
        if (!newCardData.title && !newCardData.description && !newCardData.contact) {
            setIsAdding(false);
            return;
        }

        try {
            const response = await axios.post('http://localhost:8080/api/v1/ticket', {
                title: newCardData.title || "",
                description: newCardData.description || "",
                status: status,
                contact: newCardData.contact || "",
            });

            onCardUpdate((prevCards) => {
                const newCards = [...prevCards, response.data];
                scrollToTop(newCards.length - 1);
                return newCards;
            });

            setError(null);
            setIsAdding(false);
            setNewCardData({ title: '', description: '', contact: '' });

        } catch (error) {
            console.error('Error adding card:', error);
            setError('Failed to add card');
        }
    };

    const scrollToTop = (index) => {
        if (cardRefs.current[index]) {
            cardRefs.current[index].scrollIntoView({ behavior: 'smooth', block: 'start' });
        }
    };

    return (
        <div ref={drop} className="kanban-column">
            <h2>{status.charAt(0).toUpperCase() + status.slice(1)} ({cards.length})</h2>


            <button onClick={() => setIsAdding(true)}>+ Add Card</button>

            {isAdding && (
                <div ref={formRef} className="add-card-form">
                    <input
                        type="text"
                        name="title"
                        placeholder="Title"
                        value={newCardData.title}
                        onChange={handleInputChange}
                        onBlur={handleAddCard}
                        onKeyDown={handleKeyDown}
                    />
                    <textarea
                        name="description"
                        placeholder="Description"
                        value={newCardData.description}
                        onChange={handleInputChange}
                        onBlur={handleAddCard}
                        onKeyDown={handleKeyDown}
                    />
                    <input
                        type="text"
                        name="contact"
                        placeholder="Contact"
                        value={newCardData.contact}
                        onChange={handleInputChange}
                        onBlur={handleAddCard}
                        onKeyDown={handleKeyDown}
                    />
                    {error && <p className="error-message">{error}</p>}
                </div>
            )}

            {cards.map((card, index) => (
                <KanbanCard
                    key={card.id}
                    card={card}
                    onCardUpdate={onCardUpdate}
                    ref={(el) => cardRefs.current[index] = el}
                />
            ))}
        </div>
    );
};

export default KanbanColumn;
