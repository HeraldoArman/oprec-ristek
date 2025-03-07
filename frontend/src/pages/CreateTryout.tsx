import React from "react";

import TryoutFormComponent from "../component/CreateTryoutForm";
import FooterComponent from "../component/FooterComponent";
import NavbarComponent from "../component/NavbarComponent";
import { Flowbite } from "flowbite-react";
const CreateTryout = () => {
  return (
    <Flowbite>
      <div className="relative flex min-h-screen flex-col justify-center bg-gray-100">
        <NavbarComponent />

        <div className="container mx-auto flex flex-grow items-center justify-center px-5 py-5">
          <div className="w-full max-w-lg">
            <TryoutFormComponent />
          </div>
        </div>

        <FooterComponent />
      </div>
    </Flowbite>
  );
};

export default CreateTryout;
