import { useState } from "react";
import axios from "axios";

export default function EmailReminder() {
  const [email, setEmail] = useState("");
  const [delay, setDelay] = useState(10);
  const [status, setStatus] = useState("");

  const handleSubmit = async () => {
    setStatus("Sending...");
    try {
      const res = await axios.post("http://localhost:8081/reminder", {
        email,
        delay,
      });
      setStatus(`✅ Reminder set! Workflow ID: ${res.data.workflow_id}`);
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        setStatus(
          `❌ Error: ${error.response?.data?.message || error.message}`
        );
      } else if (error instanceof Error) {
        setStatus(`❌ Error: ${error.message}`);
      } else {
        setStatus(`❌ An unknown error occurred`);
      }
    }
  };

  return (
    <div className="flex w-full h-screen items-center justify-center">
      <div className="flex flex-col items-center justify-center p-10 gap-4 rounded-2xl shadow-lg">
        <h2>Email Reminder</h2>
        <input
          type="email"
          placeholder="Your email"
          value={email}
          className="w-64 p-2 border rounded"
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="number"
          placeholder="Delay in seconds"
          value={delay}
          className="w-64 p-2 border rounded"
          onChange={(e) => setDelay(Number(e.target.value))}
        />
        <button className="w-64 p-2 border rounded hover:bg-amber-50" onClick={handleSubmit}>Set Reminder</button>
        <p>{status}</p>
      </div>
    </div>
  );
}
