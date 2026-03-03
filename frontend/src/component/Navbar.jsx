// src/components/Navbar.jsx
import { useState } from 'react';

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <nav className="backdrop-black-lg text-black shadow-lg sticky top-0 z-50">
      <div className="container mx-auto px-6 py-3 flex justify-between items-center">
        {/* Logo */}
        <div className="flex items-center">
          <div className="p-2 rounded-md mr-3">
            <img src="/Logo2.png" alt="Logo" className="h-14 w-auto" />
          </div>
        </div>

        {/* Desktop Menu (hidden di mobile) */}
        <div className="hidden md:flex space-x-8">
          <a href="#coverage" className="hover:text-yellow-400 transition">Coverage</a>
          <a href="#harga" className="hover:text-yellow-400 transition">Paket</a>
          <a href="#contact" className="hover:text-yellow-400 transition">Kontak</a>
        </div>

        {/* Hamburger Button (hanya muncul di mobile) */}
        <button 
          className="md:hidden focus:outline-none"
          onClick={toggleMenu}
          aria-label="Toggle menu"
        >
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path 
              strokeLinecap="round" 
              strokeLinejoin="round" 
              strokeWidth={2} 
              d={isMenuOpen ? "M6 18L18 6M6 6l12 12" : "M4 6h16M4 12h16M4 18h16"} 
            />
          </svg>
        </button>
      </div>

      {/* Mobile Menu (muncul saat diklik) */}
      <div className={`md:hidden ${isMenuOpen ? 'block' : 'hidden'} bg-blue-800 px-6 py-3`}>
        <div className="flex flex-col space-y-4">
          <a 
            href="#coverage" 
            className="hover:text-yellow-400 transition"
            onClick={() => setIsMenuOpen(false)}
          >
            Coverage
          </a>
          <a 
            href="#harga" 
            className="hover:text-yellow-400 transition"
            onClick={() => setIsMenuOpen(false)}
          >
            Paket
          </a>
          <a 
            href="#contact" 
            className="hover:text-yellow-400 transition"
            onClick={() => setIsMenuOpen(false)}
          >
            Kontak
          </a>
        </div>
      </div>
    </nav>
  );
}