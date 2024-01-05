import React from "react";
import { AppBar, Toolbar, Typography, Box, Tabs, Tab, IconButton, Avatar } from "@mui/material";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";

// Define the props that Header will accept
interface HeaderProps {
  balance: number;
  username: string;
  currentTab: number;
  onTabChange: (event: React.SyntheticEvent, newValue: number) => void;
}

const Header: React.FC<HeaderProps> = ({ balance, username, currentTab, onTabChange }) => {
    return (
        <AppBar position="static">
            <Toolbar>
                {/* You can place a logo here */}
                <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                    Punt
                </Typography>

                {/* Navigation Tabs */}
                <Tabs value={currentTab} onChange={onTabChange} textColor="inherit">
                    <Tab label="CURRENT GAMES" />
                    <Tab label="MY BETS" />
                    <Tab label="FRIENDS & GROUPS" />
                    {/* Add more tabs if needed */}
                </Tabs>

                {/* Balance and User Profile */}
                <Box sx={{ display: 'flex', alignItems: 'center', marginLeft: 'auto' }}>
                    <Typography variant="subtitle1" sx={{ margin: '0 20px' }}>
                        Balance: ${balance.toFixed(2)}
                    </Typography>

                    <IconButton sx={{ p: 0 }}>
                        <Avatar>
                            <AccountCircleIcon />
                        </Avatar>
                    </IconButton>
                    <Typography variant="subtitle1" sx={{ marginLeft: '10px' }}>
                        {username}
                    </Typography>
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default Header;
