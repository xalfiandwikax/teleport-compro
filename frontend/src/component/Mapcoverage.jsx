// components/MapCoverage.jsx
import IndonesiaMap from "/indonesia.svg";

const cities = [
  { name: "Jakarta", top: "58%", left: "35%" },
  { name: "Surabaya", top: "63%", left: "55%" },
  { name: "Medan", top: "48%", left: "28%" },
  { name: "Makassar", top: "63%", left: "70%" },
  { name: "Denpasar", top: "68%", left: "60%" },
  { name: "Balikpapan", top: "60%", left: "65%" },
  { name: "Jayapura", top: "58%", left: "90%" },
  { name: "Palembang", top: "58%", left: "40%" },
];

export default function MapCoverage() {
  return (
    <section className="py-20 bg-gray-50">
      <div className="max-w-7xl mx-auto px-6">

        <div className="grid lg:grid-cols-2 gap-16 items-center">

          {/* LEFT CONTENT */}
          <div>

            <p className="text-blue-600 font-semibold mb-3">
              COVERAGE AREA
            </p>

            <h2 className="text-4xl font-bold mb-6">
              JANGKAUAN <span className="text-blue-600">TERLUAS</span>
            </h2>

            <p className="text-gray-600 mb-8 max-w-md">
              Teleport.Net hadir di titik-titik strategis Indonesia untuk 
              mendukung akselerasi bisnis Anda melalui infrastruktur 
              backbone yang handal.
            </p>

            {/* City List */}
            <div className="grid grid-cols-2 gap-4">
              {cities.map((city, index) => (
                <div
                  key={index}
                  className="flex items-center gap-3 bg-white border border-gray-200 rounded-lg px-4 py-3 shadow-sm hover:shadow-md transition"
                >
                  <div className="w-3 h-3 bg-blue-500 rounded-full"></div>
                  <span className="text-gray-700 font-medium">
                    {city.name}
                  </span>
                </div>
              ))}
            </div>

          </div>

          {/* RIGHT MAP */}
          <div className="relative">

            {/* SLA Badge */}
            <div className="absolute -top-6 -right-6 bg-blue-600 text-white px-6 py-4 rounded-xl shadow-lg rotate-3 z-10">
              <p className="text-2xl font-bold">98%</p>
              <p className="text-xs">SLA UPTIME</p>
            </div>

            <div className="bg-white rounded-3xl shadow-xl p-6 relative">

              {/* SVG MAP */}
              <img
                src={IndonesiaMap}
                alt="Indonesia Map"
                className="w-full"
              />

              {/* MARKERS */}
              {cities.map((city, index) => (
                <div
                  key={index}
                  className="absolute"
                  style={{ top: city.top, left: city.left }}
                >
                  <div className="relative">

                    {/* Dot */}
                    <div className="w-3 h-3 bg-blue-600 rounded-full"></div>

                    {/* Ping Animation */}
                    <div className="absolute inset-0 w-3 h-3 bg-blue-400 rounded-full animate-ping"></div>

                  </div>
                </div>
              ))}

            </div>

          </div>

        </div>

      </div>
    </section>
  );
}