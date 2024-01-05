import React from "react";
import { Typography } from "@mui/material";

interface BalanceDisplayProps {
    balance: number; 
}

const BalanceDisplay: React.FC<BalanceDisplayProps> = ({ balance }) => {
    return <Typography>Balance: ${balance.toFixed(2)}</Typography>;
}; 

export default BalanceDisplay; 