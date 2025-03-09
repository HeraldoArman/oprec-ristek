import { useEffect, useState } from "react";
import CardComponent from "./CardComponent";
import Search from "./SearchComponent";
import { Spinner } from "flowbite-react";
import { useNavigate } from "react-router-dom";

interface Tryout {
  id: number;
  title: string;
  detail: string;
  image?: string;
  kategori: string;
}

interface BoxComponentProps {
  name: string;
  addTryout: boolean;
  tryouts: Tryout[];
  loading: boolean;
  error: string;
  setTryouts: (id: number) => void;
  searchTerm: string;
  setSearchTerm: (term: string) => void;
}

const BoxComponent: React.FC<BoxComponentProps> = ({
  name,
  addTryout,
  tryouts,
  loading,
  error,
  setTryouts,
  searchTerm,
  setSearchTerm,
}) => {
  const [filteredTryouts, setFilteredTryouts] = useState<Tryout[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    console.log("box", searchTerm);
    setFilteredTryouts(tryouts);
  }, [tryouts]);

  return (
    <div className="mx-auto max-w-4xl rounded-lg bg-white p-6 shadow-md">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold">{name}</h2>
        <Search searchTerm={searchTerm} setsearchTerm={setSearchTerm} />
      </div>

      {loading && (
        <div className="flex justify-center">
          <Spinner color="success" aria-label="spinner" />
        </div>
      )}
      {error && <p className="text-center text-red-500">{error}</p>}

      {!loading && !error && filteredTryouts.length > 0 ? (
        <div className="mt-4 grid grid-cols-1 gap-4 md:grid-cols-3">
          {filteredTryouts.map((course) => (
            <CardComponent
              key={course.id}
              {...course}
              addTryout={addTryout}
              setTryouts={setTryouts}
            />
          ))}
        </div>
      ) : (
        !loading &&
        !error && (
          <p className="text-center text-gray-500">No tryouts available.</p>
        )
      )}

      {addTryout && (
        <div className="flex justify-end py-3 pt-5">
          <button
            className="rounded-lg bg-green-500 px-6 py-2 font-semibold text-white hover:bg-green-600"
            onClick={() => navigate("/create")}
          >
            Add New Tryout
          </button>
        </div>
      )}
    </div>
  );
};

export default BoxComponent;
