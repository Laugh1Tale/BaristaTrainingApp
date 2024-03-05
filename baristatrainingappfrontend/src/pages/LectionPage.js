import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { InformationPage } from './InformationPage'
import '../index.css';

export const LectionPage = () => {
    const [lections, setLection] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/lections');
          setLection(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
        <div>
        <h1>Lections</h1>
        <ul>
          {lections.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                              <Link to="/information">
                                <h3>{item.theme}</h3>
                              </Link>
              
                              <Routes>
                              <Route path="/information" element={<InformationPage />} />
                              </Routes>
                              </div> 
            ) : <h3>{item.theme}</h3>}
            </li>
          ))}
        </ul>
      </div>
    )
}