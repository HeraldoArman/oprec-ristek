// import React from "react";

import DetailComponent from "../component/DetailComponent";
import FooterComponent from "../component/FooterComponent";
import NavbarComponent from "../component/NavbarComponent";
import { Flowbite } from "flowbite-react";
const TryoutDetail = () => {
  return (
    <Flowbite>
      <div className="relative flex min-h-screen flex-col justify-center bg-gray-100">
        <NavbarComponent />

        <div className="container mx-auto flex flex-grow items-center justify-center px-2 py-2">
          <div className="w-full max-w-3xl">
            <DetailComponent />
          </div>
        </div>

        <FooterComponent />
      </div>
    </Flowbite>
  );
};

export default TryoutDetail;
