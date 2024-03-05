import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { SignUp } from './SignUp'
import { CoursesPage } from "./CoursesPage";

export const SignIn = () => {
    return (
        <div className="row">
            <div className="col s6 offset-s3">
                <h1>Barista Training App</h1>
                <div className="card blue darken-1">
        <div className="card-content white-text">
          <span className="card-title">Авторизация</span>
          
          <div>
          <div className="input-field">
            <input id="email_inline" type="email" className="validate" />
            <label for="email_inline">Email</label>
            <span className="helper-text" data-error="wrong" data-success="right">Helper text</span>
          </div>
    
    <div className="row">
        <div className="input-field">
          <input id="password" type="password" classNames="validate" />
          <label for="password">Password</label>
        </div>
      </div>

          </div>
        </div>
        <div className="card-action">
            <Link to="/courses">
            <button className="btn blue-darken-3" style={{marginRight: 10}}>Войти</button>
            </Link>

            <Routes>
            <Route path="/courses" element={<CoursesPage />} />
            </Routes>
            <Link to="/sign-up">
            <button className="btn grey lighten-1 black-text">Регистрация</button>
            </Link>

            <Routes>
              <Route path="/sign-up" element={<SignUp />} />
            </Routes>
        </div>
      </div>
            </div>
        </div>
    )
}