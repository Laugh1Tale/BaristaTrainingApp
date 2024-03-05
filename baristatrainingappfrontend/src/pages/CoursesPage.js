import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LessonsPage } from './LessonsPage'

export const CoursesPage = () => {
    const [courses, setCourse] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/courses'); // Замените 'YOUR_API_ENDPOINT' на ваш адрес API
          setCourse(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
        <div>
        <h1>Courses</h1>
        <ul>
          {courses.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                              <Link to="/lessons">
                                <h3>{item.name}</h3>
                              </Link>
              
                              <Routes>
                              <Route path="/lessons" element={<LessonsPage />} />
                              </Routes>
                              </div> // Замените на свой URL для первой новости
            ) : <h3>{item.name}</h3>}
              <p>{item.description}</p>
              <p>Проходной балл</p>
              <p>{item.passing_score}</p>
            </li>
          ))}
        </ul>
      </div>
    )
}