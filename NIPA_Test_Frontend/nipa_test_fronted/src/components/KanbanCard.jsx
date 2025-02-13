//KanbanCard.jsx

import React, { useState, useRef, useEffect } from 'react';
import { useDrag } from 'react-dnd';
import axios from 'axios';
import '../styles/KanbanCard.css';

const KanbanCard = ({ card, onCardUpdate }) => {

    
    const [{ isDragging }, drag] = useDrag({
        type: 'CARD',
        item: { id: card.id, status: card.status },
        collect: (monitor) => ({
            isDragging: monitor.isDragging(),
        }),
    });

    const [editedData, setEditedData] = useState({
        title: card.title,
        description: card.description,
        contact: card.contact
    });

    const descriptionRef = useRef(null);

    const handleUpdate = async (field, value) => {
        const updatedCard = { 
            ...card,
            [field]: value
        };

        setEditedData((prev) => ({ ...prev, [field]: value }));

        try {
            const response = await axios.put(`http://localhost:8080/api/v1/ticket/${card.id}`, updatedCard);
            if (response.data) {
                onCardUpdate(card.id, updatedCard); 
            }
        } catch (error) {
            console.error("Error updating card:", error);
        }
    };
    const handleDescriptionChange = (e) => {
        setEditedData({ ...editedData, description: e.target.value });
    };

    useEffect(() => {
        if (descriptionRef.current) {
            const length = descriptionRef.current.value.length;
            descriptionRef.current.setSelectionRange(length, length);
        }
    }, [editedData.description]);

    return (
        <div ref={drag} className="kanban-card" style={{ opacity: isDragging ? 0.5 : 1 }}>
            <input
                type="text"
                value={editedData.title}
                onChange={(e) => setEditedData({ ...editedData, title: e.target.value })}
                onBlur={(e) => handleUpdate('title', e.target.value)}
                placeholder="Title"
            />
            <textarea
                ref={descriptionRef}
                value={editedData.description}
                onChange={handleDescriptionChange}
                onBlur={(e) => handleUpdate('description', e.target.value)}
                placeholder="Description"
            />
            <input
                type="text"
                value={editedData.contact}
                onChange={(e) => setEditedData({ ...editedData, contact: e.target.value })}
                onBlur={(e) => handleUpdate('contact', e.target.value)}
                placeholder="Contact"
            />
            <p>Updated:{card.updated ? new Date(card.updated).toLocaleString() : 'Not updated yet'}</p>
        </div>
    );
};

export default KanbanCard;
