import React, { useState } from 'react';
import Header from '../components/Header';
import BalanceDisplay from '../components/BalanceDisplay';
import NavigationTabs from '../components/NavigationTabs';
import { Container } from '@mui/material';
import { Box, Card, CardContent } from '@mui/material';
import { useNavigate } from 'react-router-dom'; // Make sure to import useNavigate hook



const Home = () => {
    const [balance, setBalance] = useState(0); // This would be fetched from your user state or context
    const [currentTab, setCurrentTab] = useState(0);
    const [username, setUsername] = useState('User'); // Replace with actual username from auth state or context
    const navigate = useNavigate();

    const handleTabChange = (
        event: React.SyntheticEvent<Element, Event>,
        newValue: number // Ensure this is number if you're using indices
      ) => {
        setCurrentTab(newValue);
    };

    return (
        <div>
            <Header 
              balance={balance} 
              username={username}
              currentTab={currentTab} 
              onTabChange={handleTabChange} 
            />
            {/* Rest of the content */}
        </div>
    );
}; 

export default Home; 
