import React, { useState } from 'react';
import Header from '../components/Header';
import BalanceDisplay from '../components/BalanceDisplay';
import NavigationTabs from '../components/NavigationTabs';
import { Container } from '@mui/material';

const Home = () => {
    const [balance, setBalance] = useState(0);
    const [currentTab, setCurrentTab] = useState(0);

    const handleTabChange = (
        event: React.SyntheticEvent<Element, Event>,
        newValue: number // Ensure this is number if you're using indices
      ) => {
        setCurrentTab(newValue);
    };


    return (
        <div>
            <Header /> 
            <Container>
                <BalanceDisplay balance={balance} />
                <NavigationTabs currentTab={currentTab} onChange={handleTabChange} />
                {currentTab === 0 && <div>Current Games Content</div>}
                {currentTab === 1 && <div>My Bets Content</div>}
                {currentTab === 2 && <div>Friends & Groups Content</div>}
            </Container>
        </div>
    ); 
}; 

export default Home; 
