import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { TestsPage } from './TestsPage'
import '../css/courses.css';

export const InformationPage = () => {
    const [information, setInformation] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/information');
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
          const response = await axios.get('http://localhost:8000/api/tests');
          setTest(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    return (
      <div class="ag-format-container">
      <div class="ag-courses_box">

        {information.map((item, index) => (

          <div class="ag-courses_item">
            <a href="/information" class="ag-courses-item_link">
            <div class="ag-courses-item_bg"></div>

            <div class="ag-courses-item_title">
              {item.theme}
            </div>

            <div class="ag-courses-item_date-box">
            </div>

            <div class="ag-courses-item_date-box">
              {"\n"}
            </div>

          <div class="ag-courses-item_date-box">
            <span class="ag-courses-item_date">
              {item.text}
            </span>
          </div>
          </a>
        </div>
        ))}

      </div>

      <section className="container forms">
          <div className="form login">
            <div className="form-content">
              <header>Test</header>
              <form action="#">
                <h2>{tests.theme}</h2>
                <h3>{tests.description}</h3>

                <Link to="/tests">
                  <button class="field button-field" style={{marginRight: 10}}>Решить тест</button>
                </Link>

                <Routes>
                  <Route path="/tests" element={<TestsPage />} />
                </Routes>
              </form>
            </div>
          </div>
        </section>
    </div>
    )
}
