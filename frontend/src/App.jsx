// src/App.jsx
import { BrowserRouter as Router } from 'react-router-dom';
import Navbar from './component/Navbar';
import HeroSection from './component/HeroSection';
import PricingTable from './component/PricingTable';
import Mapcoverage from './component/Mapcoverage';
import ContactISP from './component/ContactISP';
import Footer from './component/Footer';
import 'leaflet/dist/leaflet.css'; 

function App() {
  return (
    <Router>
      <div className="font-sans bg-gray-50 min-h-screen">
        <Navbar />
        <HeroSection 
          title="Internet Service Provider untuk Rumah dan Bisnis Anda" 
          subtitle="TELEPORT.NET - Provider Fiber Optic Tercepat di Indonesia"
          ctaText="CEK KETERSEDIAAN"
        />
        <Mapcoverage />
        <PricingTable />
        <ContactISP />
        <Footer 
          companyName="TELEPORT.NET"
          address="Jl. Gatot Subroto No. 42, Jakarta Selatan"
          phone="+62 21 1234 5678"
          email="cs@teleport.net"
        />
      </div>
    </Router>
  );
}

export default App;