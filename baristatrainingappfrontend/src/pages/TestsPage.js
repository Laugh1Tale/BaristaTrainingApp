import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LessonsPage } from './LessonsPage'
import '../css/tests.css';

export const TestsPage = () => {
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/questions');
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
          const response = await axios.get('http://localhost:8000/api/answers');
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
    <div className='radio-item'>
      <input name="radio" id="radio1" type="radio" />
      <label for="radio1">{answers[0].text}</label>
    </div>

    <div className='radio-item'>
      <input name="radio" id="radio2" type="radio" />
      <label for="radio2">{answers[1].text}</label>
    </div>
  </form>
                              </div> 
            ) : <div></div>}
               {index === 1 ? (
                              <div>
                                <h2>{item.theme}</h2>
                                <h3>{item.text}</h3>
                                <form action="#">
                                <div className='radio-item'>
      <input name="radio" id="radio3" type="radio" />
      <label for="radio3">{answers[2].text}</label>
    </div>

    <div className='radio-item'>
      <input name="radio" id="radio4" type="radio" />
      <label for="radio4">{answers[3].text}</label>
    </div>
  </form>
                              </div>
            ) : <div></div>}
               {index === 2 ? (
                            <div>
                                <h2>{item.theme}</h2>
                                <h3>{item.text}</h3>
                                <form action="#">
    <div className='radio-item'>
      <input name="radio" id="radio5" type="radio" />
      <label for="radio5">{answers[4].text}</label>
    </div>

    <div className='radio-item'>
      <input name="radio" id="radio6" type="radio" />
      <label for="radio6">{answers[5].text}</label>
    </div>
  </form>
                            </div>
            ) : <div></div>}
            </li>
          ))}
        </ul>
      </div>

      <section className="container forms">
          <div className="form login">
            <div className="form-content">
              <form action="#">
                <Link to="/lessons">
                  <button class="field button-field" >Закончить тест</button>
                </Link>

                <Routes>
                  <Route path="/lessons" element={<LessonsPage />} />
                </Routes>
              </form>
            </div>
          </div>
        </section>
      </div>
    )
}