import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { TestsPage } from './TestsPage'

export const InformationPage = () => {
    const [information, setInformation] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/information'); // Замените 'YOUR_API_ENDPOINT' на ваш адрес API
          setInformation(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    const [test, setTest] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/tests'); // Замените 'YOUR_API_ENDPOINT' на ваш адрес API
          setInformation(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
        <div>
        <h1>information</h1>
        <ul>
          {information.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                              <Link to="/tests">
                                <h3>{item.name}</h3>
                              </Link>
              
                              <Routes>
                              <Route path="/tests" element={<TestsPage />} />
                              </Routes>
                              </div> // Замените на свой URL для первой новости
            ) : <h3>{item.name}</h3>}
              <p>{item.theme}</p>
              <p>{item.text}</p>
            </li>
          ))}
        </ul>
        <h1>Test</h1>   
        <h2>{test.theme}</h2>
        <h3>{test.description}</h3>
      </div>
    )
}