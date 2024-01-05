import React from "react";
import { Tabs, Tab } from '@mui/material'

interface NavigationTabsProps {
    currentTab: number; 
    onChange: (event: React.SyntheticEvent, newValue: number) => void;
}

const NavigationTabs: React.FC<NavigationTabsProps> = ({ currentTab, onChange }) => {
    return (
        <Tabs value={currentTab} onChange={onChange}>
            <Tab label="Current Games" />
            <Tab label="My Bets" />
            <Tab label="Friends & Groups" />
        </Tabs>
    ); 
};

export default NavigationTabs

