import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { Card } from "flowbite-react";
import { Spinner } from "flowbite-react";

const DetailComponent = () => {
  const { id } = useParams();
  const [formData, setFormData] = useState({
    title: "",
    detail: "",
    image: "",
    username: "john_doe3",
    createdAt: "",
    kategori: "",
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchTryout = async () => {
      setLoading(true);
      try {
        const response = await fetch(`http://127.0.0.1:3000/tryout/${id}`);
        if (!response.ok) throw new Error("Gagal mengambil data");
        const data = await response.json();
        const date = new Date(data.CreatedAt);
        const formattedDate = date.toLocaleDateString("en-GB", {
          day: "2-digit",
          month: "long",
          year: "numeric",
        });
        setFormData({
          title: data.title,
          detail: data.detail,
          image: data.image || "https://picsum.photos/1000/600",
          username: data.username,
          kategori: data.kategori,
          createdAt: formattedDate,
        });
      } catch (error) {
        console.error("Error:", error);
        alert("Terjadi kesalahan saat mengambil data tryout.");
      } finally {
        setLoading(false);
      }
    };

    if (id) fetchTryout();
  }, [id]);
  return (
    <div className="relative w-full">
      <Card className="relative w-full">
        {loading ? (
          <div className="flex justify-center">
            <Spinner color="success" aria-label="spinner" />
          </div>
        ) : (
          <img
            className="h-64 w-full rounded-lg object-cover"
            src={formData.image}
            alt={formData.image}
          />
        )}
        <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
          {formData.title}
        </h5>
        <p className="font-normal text-gray-500 dark:text-gray-400">
          Date Created: {formData.createdAt}
        </p>
        <p className="font-normal text-gray-700 dark:text-gray-400">
          Category: {formData.kategori}
        </p>
        <p className="font-normal text-gray-700 dark:text-gray-400">
          {formData.detail}
        </p>
        <div className="flex justify-end py-3 pt-5">
          <button className="rounded-lg bg-green-500 px-6 py-2 font-semibold text-white hover:bg-green-600">
            Start Now (Coming soon)
          </button>
        </div>
      </Card>
    </div>
  );
};

export default DetailComponent;
