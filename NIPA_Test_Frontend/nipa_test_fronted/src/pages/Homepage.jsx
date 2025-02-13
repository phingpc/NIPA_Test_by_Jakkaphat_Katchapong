// Homepage.jsx

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import KanbanColumn from '../components/KanbanColumn.jsx';
import '../styles/Homepage.css';

const Homepage = () => {
    const [cards, setCards] = useState({
        pending: [],
        accepted: [],
        resolved: [],
        rejected: []
    });
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchCards = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/v1/ticket');
                let data = response.data;

                if (!Array.isArray(data)) {
                    data = [data];
                }

                setCards({
                    pending: data.filter(card => card.status === 'pending'),
                    accepted: data.filter(card => card.status === 'accepted'),
                    resolved: data.filter(card => card.status === 'resolved'),
                    rejected: data.filter(card => card.status === 'rejected')
                });

                setError(null);
            } catch (error) {
                console.error('Error fetching cards:', error);
                setError('Failed to load tickets');
            } finally {
                setLoading(false);
            }
        };

        fetchCards();
    }, [cards]);

    const handleCardUpdate = (cardId, updatedData) => {
        setCards(prevCards => {
            const updatedCards = { ...prevCards };

            Object.keys(updatedCards).forEach(status => {
                updatedCards[status] = updatedCards[status].map(card =>
                    card.id === cardId ? { ...card, ...updatedData } : card
                );
            });

            return updatedCards;
        });
    };




    const handleDrop = async (cardId, newStatus) => {
        try {
            console.log("Moving card:", cardId, "to", newStatus);

            const cardToUpdate = Object.values(cards).flat().find(card => card.id === cardId);

            if (!cardToUpdate || cardToUpdate.status === newStatus) {
                console.log("ลากไปที่เดิม ไม่ต้องอัปเดต");
                return;
            }

            const updatedCard = {
                ...cardToUpdate,
                status: newStatus,
                updated: new Date().toISOString()
            };

            console.log("Updating card:", updatedCard);

            await axios.put(`http://localhost:8080/api/v1/ticket/${cardId}`, updatedCard);

            setCards(prevCards => {
                let updatedCards = { ...prevCards };
                let movedCard = null;

                Object.keys(updatedCards).forEach(status => {
                    updatedCards[status] = updatedCards[status].filter(card => {
                        if (card.id === cardId) {
                            movedCard = { ...updatedCard };
                            return false;
                        }
                        return true;
                    });
                });

                
                if (movedCard) {
                    updatedCards[newStatus] = updatedCards[newStatus] || [];
                    updatedCards[newStatus].push(movedCard);
                }

                console.log("Updated cards:", updatedCards);
                return updatedCards;
            });

        } catch (error) {
            console.error("Error updating card status:", error);
        }
    };



    return (
        <DndProvider backend={HTML5Backend}>
            <div className='container'>
                {loading && <p>Loading...</p>}
                <div className="kanban-board">
                    {['pending', 'accepted', 'resolved', 'rejected'].map(status => (
                        <KanbanColumn
                            key={status}
                            status={status}
                            cards={cards[status]}
                            onDrop={handleDrop}
                            onCardUpdate={handleCardUpdate}
                        />
                    ))}
                </div>
            </div>
            <div className='container-error'>
            {loading && <div className='error-container'><p>Loading...</p></div>}
            {error && <div className='error-container'><p>{error}</p></div>}
            </div>
        </DndProvider>

    );
};

export default Homepage;