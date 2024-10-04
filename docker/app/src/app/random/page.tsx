"use client";

import React, { useEffect, useState } from "react";
async function getRandom() {
  // TODO: change this
  const res = await fetch(`http://localhost:8080/random`);
  return res.text();
}

export default async function Home() {
  const [num, setNum] = useState<any | null>();

  useEffect(() => {
    async function fetchNumber() {
      const number = await getRandom();
      const nu = await Promise.all([number]);
      setNum(nu);
    }

    fetchNumber();
    console.log(num);
  }, []);

  return (
    <div className="fancy-number-container">
      <span className="fancy-number">{num}</span>
    </div>
  );
}
