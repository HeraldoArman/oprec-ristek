import { Card } from "flowbite-react";
import CardComponent from "./CardComponent";
import { useState } from "react";
import Search from "./SearchComponent";
const BoxComponent = ({
  name,
  tryout,
  addTryout,
}: {
  name: string;
  tryout: {
    title: string;
    description: string;
    image: string;
    buttonLabel: string;
    buttonStyle: string;
  }[];
  addTryout: boolean;
}) => {
  const [searchTerm, setsearchTerm] = useState("");

  return (
    <div className="mx-auto max-w-4xl rounded-lg bg-white p-6 shadow-md">
      <div className="flex items-center justify-between">
        <h2 className="font-bold text-2xl">{name}</h2>
        {!addTryout ? (
            <Search searchTerm={searchTerm} setsearchTerm={setsearchTerm} />
        ): (
        <a href="#" className="text-sm font-medium text-green-500">
          View All
        </a>
        )}
      </div>

      <div className="mt-4 grid grid-cols-1 gap-4 md:grid-cols-3">
        {tryout.map((course, index) => (
          <CardComponent key={index} {...course} />
        ))}
      </div>
      {addTryout ? (
      <div className="flex justify-end py-3 pt-5">
        <button className="rounded-lg bg-green-500 px-6 py-2 font-semibold text-white hover:bg-green-600">
          Add New Tryout
        </button>
      </div>

      ) : (<></>)}
    </div>
  );
};
export default BoxComponent;
