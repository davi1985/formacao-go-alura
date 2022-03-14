import reactLogo from "../../assets/react.svg";
import goLangLogo from "../../assets/go.png";

import "./styles.css";

export const Header = () => (
  <header className="header">
    <img src={goLangLogo} width={50} className="logo" alt="logo" />

    <h1>Go and React Application</h1>

    <img src={reactLogo} width={50} className="logo" alt="logo" />
  </header>
);
