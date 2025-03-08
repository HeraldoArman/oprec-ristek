import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, Label, TextInput, Textarea, Spinner } from "flowbite-react";

const TryoutFormComponent = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    title: "",
    detail: "",
    image : "",
    username: "john_doe3",
  });
  const [loading, setLoading] = useState(false); 

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({
      ...formData,
      [e.target.id]: e.target.value,
    });
  };

  // Handle submit form
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true); // Aktifkan loading

    try {
      const response = await fetch("http://127.0.0.1:3000/tryout", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error("Gagal membuat tryout");
      }

      setFormData({ title: "",image:"", detail: "", username: "john_doe3" });
      navigate("/");
    } catch (error) {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat membuat tryout.");
    } finally {
      setLoading(false); 
    }
  };

  return (
    <div className="relative w-full max-w-lg">
      <Card className="relative w-full">
        {loading && (
          <div className="absolute inset-0 bg-white bg-opacity-70 flex items-center justify-center rounded-lg">
            <Spinner size="xl" />
          </div>
        )}

        <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="title" value="Title" />
            </div>
            <TextInput
              id="title"
              type="text"
              placeholder="Tryout Name"
              required
              value={formData.title}
              onChange={handleChange}
              disabled={loading} // 
            />
          </div>

          <div>
            <div className="mb-2 block">
              <Label htmlFor="image" value="Image Link" />
            </div>
            <TextInput
              id="image"
              type="url"
              placeholder="Image link (optional)"

              value={formData.image}
              onChange={handleChange}
              disabled={loading} // 
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="detail" value="Description" />
            </div>
            <Textarea
              id="detail"
              placeholder="Tryout Description"
              required
              rows={5}
              className="resize-none"
              value={formData.detail}
              onChange={handleChange}
              disabled={loading}
            />
          </div>

          <div className="flex justify-end py-3 pt-5">
            <button
              type="submit"
              className="rounded-lg bg-green-500 px-6 py-2 font-semibold text-white hover:bg-green-600 disabled:opacity-50"
              disabled={loading} 
            >
              {loading ? "Processing..." : "Create Tryout"}
            </button>
          </div>
        </form>
      </Card>
    </div>
  );
};

export default TryoutFormComponent;
