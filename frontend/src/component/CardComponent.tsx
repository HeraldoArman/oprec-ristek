import { Card, Dropdown, Spinner } from "flowbite-react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

interface CardComponentProps {
  id: number;
  title: string;
  detail: string;
  image?: string;
  addTryout?: boolean;
  kategori: string;
  setTryouts: (id: number) => void;
}

const CardComponent: React.FC<CardComponentProps> = ({
  id,
  title,
  detail,
  image = "https://via.placeholder.com/150",
  addTryout = false,
  kategori,
  setTryouts,
}) => {
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const handleDelete = async () => {
    try {
      const response = await fetch(`http://127.0.0.1:3000/tryout/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      });

      if (!response.ok) {
        throw new Error("Gagal menghapus tryout");
      }

      setTryouts(id);
    } catch (error) {
      console.error("Error deleting tryout:", error);
      alert("Terjadi kesalahan saat menghapus tryout.");
    }
  };

  return (
    <Card className="relative w-full p-4 shadow-md">
      {addTryout && (
        <div className="absolute right-2 top-2">
          <Dropdown inline label="">
            <Dropdown.Item>
              <a
                href={`/edit/${id}`}
                className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              >
                Edit
              </a>
            </Dropdown.Item>
            <Dropdown.Item onClick={handleDelete}>
              <span className="block px-4 py-2 text-sm text-red-600 hover:bg-gray-100">
                Delete
              </span>
            </Dropdown.Item>
          </Dropdown>
        </div>
      )}

      {loading && (
        <div className="flex justify-center">
          <Spinner color="success" aria-label="spinner" />
        </div>
      )}

      <img
        src={image}
        alt={image}
        className={`h-40 w-full rounded-t-lg object-cover ${loading ? "hidden" : "block"}`}
        onLoad={() => setLoading(false)}
        onError={() => setLoading(false)}
      />

      <div className="p-4">
        <h3 className="text-center text-lg font-semibold">{title}</h3>
        <p className="pb-2 text-center text-xs text-gray-600">
          category: {kategori}
        </p>
        <p className="pb-2 text-center text-xs text-gray-600">{detail}</p>
        <button
          className="w-full rounded-lg bg-blue-500 py-2 text-white hover:bg-blue-600"
          onClick={() => navigate(`/detail/${id}`)}
        >
          Start Tryout
        </button>
      </div>
    </Card>
  );
};

export default CardComponent;
