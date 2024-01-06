import React, { createContext, useState, useContext, ReactNode } from 'react';
import { useNavigate } from 'react-router-dom';

interface AuthContextType {
  isAuthenticated: boolean;
  login: (email: string, password: string) => void;
  logout: () => void;
}

// Create the context with a default value
const AuthContext = createContext<AuthContextType>(null!);

// Define the props for the provider component
interface AuthProviderProps {
  children: ReactNode;
}

// Create the provider component
export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(
    localStorage.getItem('isAuthenticated') === 'true'
  );

  const navigate = useNavigate(); 

  const login = async (email: string, password: string) => {
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          password: password,
        }),
      });
  
      if (!response.ok) {
        // Handle HTTP errors, e.g., 401 Unauthorized
        throw new Error(`HTTP error: ${response.status}`);
      }
  
      // Check if the response has content
      const text = await response.text(); // First read it as text
      let data;
      if (text) {
        // If there's content, try to parse it as JSON
        data = JSON.parse(text);
        console.log("Login successful", data);
        // Handle successful login here, e.g., store token if provided
      } else {
        // If there's no content, proceed without parsing
        console.log("Login successful, no response data");
        // Handle successful login with no response data here
      }
      setIsAuthenticated(true);
      localStorage.setItem("isAuthenticated", "true"); 
      navigate('/home'); // Redirect to home page on successful login
    } catch (error) {
      console.error("Login failed", error);
      // Handle login failure here, e.g., show error message
    }
  };

  const logout = () => {
    // Perform the logout logic (e.g., remove the token from the secure storage)
    setIsAuthenticated(false);
  };

  // The value prop of the provider will provide these to any components that consume the context
  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// Export the useAuth hook for easy access to the context
export const useAuth = () => useContext(AuthContext);
