
CREATE TYPE ticket_status AS ENUM ('pending', 'accepted', 'resolved', 'rejected');

CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    contact VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ticket_status DEFAULT 'pending'
);


CREATE OR REPLACE FUNCTION update_ticket_timestamp() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_update_ticket
BEFORE UPDATE ON tickets
FOR EACH ROW
EXECUTE FUNCTION update_ticket_timestamp();


-- เพิ่มข้อมูลตัวอย่าง
INSERT INTO tickets (title, description, contact, status) VALUES
('Issue with login', 'Unable to login with correct credentials', 'user1@example.com', 'pending'),
('Page not loading', 'The homepage is not loading properly', 'user2@example.com', 'accepted'),
('Error in payment', 'Payment gateway is throwing an error', 'user3@example.com', 'resolved'),
('Feature request', 'Request to add dark mode feature', 'user4@example.com', 'rejected');