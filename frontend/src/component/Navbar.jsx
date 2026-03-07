// src/components/Navbar.jsx
import { useState } from "react";

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const menu = [
    { name: "Home", link: "#" },
    { name: "About", link: "#about" },
    { name: "Services", link: "#services" },
    { name: "Coverage", link: "#coverage" },
  ];

  return (
    <nav className="bg-white/80 backdrop-blur-md shadow-sm sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-6 py-3 flex items-center justify-between">
        
        {/* Logo */}
        <div className="flex items-center gap-3">
          <img src="/Logo2.png" alt="Teleport" className="h-10 w-auto" />
        </div>

        {/* Desktop Menu */}
        <div className="hidden md:flex items-center gap-10 text-gray-600 font-medium">
          {menu.map((item, index) => (
            <a
              key={index}
              href={item.link}
              className="hover:text-blue-600 transition"
            >
              {item.name}
            </a>
          ))}
        </div>

        {/* Right Button */}
        <div className="hidden md:block">
          <button className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-full shadow-md transition">
            Hubungi Kami
          </button>
        </div>

        {/* Mobile Button */}
        <button
          className="md:hidden"
          onClick={() => setIsMenuOpen(!isMenuOpen)}
        >
          <svg
            className="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d={
                isMenuOpen
                  ? "M6 18L18 6M6 6l12 12"
                  : "M4 6h16M4 12h16M4 18h16"
              }
            />
          </svg>
        </button>
      </div>

      {/* Mobile Menu */}
      <div
        className={`md:hidden bg-white border-t ${
          isMenuOpen ? "block" : "hidden"
        }`}
      >
        <div className="flex flex-col px-6 py-4 space-y-4 text-gray-600">
          {menu.map((item, index) => (
            <a
              key={index}
              href={item.link}
              onClick={() => setIsMenuOpen(false)}
              className="hover:text-blue-600"
            >
              {item.name}
            </a>
          ))}

          <button className="bg-blue-600 text-white py-2 rounded-lg">
            Hubungi Kami
          </button>
        </div>
      </div>
    </nav>
  );
}