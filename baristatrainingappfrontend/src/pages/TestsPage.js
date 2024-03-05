import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LessonsPage } from './LessonsPage'

export const TestsPage = () => {
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/questions');
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
          const response = await axios.get('http://localhost:8000/answers');
          setAnswers(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);


    return (
        <div>
        <div>
        <h1>Questions</h1>
        <ul>
          {questions.map((item, index) => (
            <li key={index}>
               {index === 0 ? (
                              <div>
                                <h2>{item.theme}</h2>
                                <h3>{item.text}</h3>
                                <form action="#">
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[0].text}</span>
      </label>
    </p>
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[1].text}</span>
      </label>
    </p>
  </form>
                              </div> 
            ) : <div></div>}
               {index === 1 ? (
                              <div>
                                <h2>{item.theme}</h2>
                                <h3>{item.text}</h3>
                                <form action="#">
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[2].text}</span>
      </label>
    </p>
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[3].text}</span>
      </label>
    </p>
  </form>
                              </div>
            ) : <div></div>}
               {index === 2 ? (
                            <div>
                                <h2>{item.theme}</h2>
                                <h3>{item.text}</h3>
                                <form action="#">
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[4].text}</span>
      </label>
    </p>
    <p>
      <label>
        <input name="group1" type="radio" />
        <span>{answers[5].text}</span>
      </label>
    </p>
  </form>
                            </div>
            ) : <div></div>}
            </li>
          ))}
        </ul>
      </div>
      <Link to="/lessons">
      <button>решить тест</button>
      </Link>

      <Routes>
      <Route path="/lessons" element={<LessonsPage />} />
      </Routes>
      </div>
    )
}