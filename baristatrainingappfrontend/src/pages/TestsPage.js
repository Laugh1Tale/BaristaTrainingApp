import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { SignUp } from './SignUp'

export const TestsPage = () => {
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/questions'); // Замените 'YOUR_API_ENDPOINT' на ваш адрес API
          setQuestions(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    const [answers, setAnswers] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/answers'); // Замените 'YOUR_API_ENDPOINT' на ваш адрес API
          setQuestions(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
        <div>
        <h1>Questions</h1>
        <ul>
          {questions.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                                <h3>{item.theme}</h3>
                              </div> // Замените на свой URL для первой новости
            ) : <div></div>}
               {index === 1 ? (
                              <div>
                                <h3>{item.theme}</h3>
                              </div> // Замените на свой URL для первой новости
            ) : <div></div>}
               {index === 2 ? (
                            <div>
                                <h3>{item.theme}</h3>
                            </div> // Замените на свой URL для первой новости
            ) : <div></div>}
            </li>
          ))}
        </ul>
      </div>
    )
}