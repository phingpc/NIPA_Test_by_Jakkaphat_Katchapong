import React from 'react';
import '../styles/footer.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGithubSquare } from '@fortawesome/free-brands-svg-icons';

const Footer = () => {
    return (
        <footer className="footer">
            <div className="footer-container-">
                <p className="footer-text">Name: jakkaphat katchapong</p>
                <p className="footer-text">Phone: 082-542-2985</p>
                <p className="footer-text">Email: katchapong_j@silpakorn.edu</p>
            </div>
            <div className="footer-container-right">
                <a href="https://github.com/phingpc" target="_blank" rel="noopener noreferrer" style={{ color: 'black'}}>
                    <FontAwesomeIcon icon={faGithubSquare} size="3x" />
                </a>
            </div>
        </footer>
    );
};

export default Footer;