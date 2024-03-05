import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { SignIn } from "./SignIn"
import { CoursesPage } from "./CoursesPage"
import '../index.css';

export const SignUp = () => {
    return (
        <div className="row">
            <div className="col s6 offset-s3">
                <h1>Barista Training App</h1>
                <div className="card blue darken-1">
        <div className="card-content white-text">
          <span className="card-title">Регистрация</span>
          
          <div>
          <div className="input-field">
            <input id="email_inline" type="email"/>
            <label for="email_inline">Email</label>
            <span className="helper-text" data-error="wrong" data-success="right">Helper text</span>
          </div>

          <div class="input-field">
          <input id="last_name" />
          <label for="last_name">Name</label>
        </div>

        <div class="input-field">
          <input id="last_name" />
          <label for="last_name">Last Name</label>
        </div>
    
    <div className="row">
        <div className="input-field">
          <input id="password" type="password"/>
          <label for="password">Password</label>
        </div>
      </div>

          </div>
        </div>
        <div className="card-action">
            <Link to="/courses">
            <button className="btn blue-darken-3" style={{marginRight: 10}}>Регистрация</button>
            </Link>

            <Routes>
            <Route path="/courses" element={<CoursesPage />} />
            </Routes>
            <Link to="/sign-in">
            <button className="btn grey lighten-1 black-text" href="./SignIn">Вход</button>
            </Link>

        <Routes>
          <Route path="/sign-in" element={<SignIn />} />
        </Routes>
        </div>
      </div>
            </div>
        </div>
    )
}