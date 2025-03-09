import React from "react";

interface SearchProps {
  searchTerm: string;
  setsearchTerm: (term: string) => void;
}

const Search: React.FC<SearchProps> = ({ searchTerm, setsearchTerm }) => {
  return (
    <div className="flex items-center rounded-full border border-gray-300 bg-white px-4 py-2 shadow-md focus-within:ring-2 focus-within:ring-green-500">
      <img
        src="/search-icon.svg"
        alt="search"
        className="h-5 w-5 text-gray-400"
      />
      <input
        type="text"
        placeholder="Search Tryout Title"
        value={searchTerm}
        onChange={(e) => setsearchTerm(e.target.value)}
        className="ml-3 w-full border-none bg-transparent text-gray-700 placeholder-gray-400 outline-none focus:ring-0"
      />
    </div>
  );
};

export default Search;
