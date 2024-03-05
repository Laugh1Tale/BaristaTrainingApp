import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LectionPage } from './LectionPage'
import { TestPoints } from './QuestionsPage'

export const LessonsPage = () => {
    const [lessons, setLesson] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/lessons');
          setLesson(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
        <div>
        <h1>Lessons</h1>
        <ul>
          {lessons.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                              <Link to="/lections">
                                <h3>{item.name}</h3>
                              </Link>
              
                              <Routes>
                              <Route path="/lections" element={<LectionPage />} />
                              </Routes>
                              </div> 
            ) : <h3>{item.name}</h3>}
              <p>{item.description}</p>
              <p>Ваш балл</p>
              {index === 0 ? <p>{TestPoints}</p>
              : <p>0</p>}
              <p>Проходной балл</p>
              <p>{item.passing_score}</p>
            </li>
          ))}
        </ul>
      </div>
    )
}