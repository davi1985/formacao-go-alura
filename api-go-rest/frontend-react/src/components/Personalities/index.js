import { useEffect, useState } from "react";
import axios from "axios";
import "./styles.css";

export const Personalities = () => {
  const [personalities, setPersonalities] = useState();

  const getData = () => {
    axios
      .get("http://localhost:8000/api/personalities")
      .then((res) => setPersonalities(res.data));
  };

  useEffect(() => {
    getData();
  }, []);

  return (
    <div className="container">
      <h1>Personalities</h1>

      {personalities?.map((personality) => (
        <div key={personality.id} className="person">
          <h3>{personality.name}</h3>
          <p>{personality.bio}</p>
        </div>
      ))}
    </div>
  );
};
