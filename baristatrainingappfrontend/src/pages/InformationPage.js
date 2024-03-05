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
          const response = await axios.get('http://localhost:8000/information');
          setInformation(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    const [tests, setTest] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/tests');
          setTest(response.data);
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
              <h2>{item.theme}</h2>
              <h3>{item.text}</h3>
            </li>
          ))}
        </ul>
        <h1>Test</h1>   
        <h2>{tests.theme}</h2>
        <h3>{tests.description}</h3>
        <div>
        <Link to="/tests">
        <button>решить тест</button>
        </Link>

        <Routes>
        <Route path="/tests" element={<TestsPage />} />
        </Routes>
        </div>
      </div>
    )
}
