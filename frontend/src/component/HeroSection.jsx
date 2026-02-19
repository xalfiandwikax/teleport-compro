import { motion } from "framer-motion";

export default function HeroSection() {
  return (
    <section className="bg-gradient-to-br from-blue-900 via-blue-800 to-blue-700 text-white">
      
      <div className="max-w-6xl mx-auto px-6 pt-28 pb-24">
        
        {/* Badge */}
        <motion.div
          initial={{ opacity: 0, y: 15 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="inline-block bg-blue-600/40 backdrop-blur px-4 py-1 rounded-full text-sm mb-8"
        >
          ISP Terbaik 2025
        </motion.div>

        {/* Title */}
        <motion.h1
          initial={{ opacity: 0, y: 40 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2, duration: 0.8 }}
          className="text-5xl md:text-6xl font-extrabold leading-tight tracking-tight"
        >
          Anytime, Anywhere <br />
          <span className="text-blue-300">
            Any Connectivity.
          </span>
        </motion.h1>

        {/* Subtitle */}
        <motion.p
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.5 }}
          className="mt-8 max-w-2xl text-lg text-blue-100 leading-relaxed"
        >
          PT. Teleport Data Persada menghadirkan internet tanpa batas wilayah.
          Akses dunia kapanpun dan dimanapun melalui Fiber Optic, Wireless,
          maupun VSAT.
        </motion.p>

        {/* Buttons */}
        <motion.div
          initial={{ opacity: 0, y: 15 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.7 }}
          className="mt-10 flex flex-wrap gap-4"
        >
          <button className="bg-blue-500 hover:bg-blue-400 px-7 py-3 rounded-lg font-semibold shadow-lg transition">
            Berlangganan Sekarang
          </button>

          <button className="border border-white/30 hover:bg-white/10 px-7 py-3 rounded-lg font-semibold transition">
            Lihat Cakupan Area
          </button>
        </motion.div>
      </div>

      {/* Statistik Box */}
      <div className="bg-white text-blue-900">
        <div className="max-w-6xl mx-auto grid grid-cols-2 md:grid-cols-4 text-center py-10 px-6">
          <div>
            <h3 className="text-2xl font-bold">98%</h3>
            <p className="text-sm text-gray-500">SLA Guarantee</p>
          </div>
          <div>
            <h3 className="text-2xl font-bold">24/7</h3>
            <p className="text-sm text-gray-500">Support Center</p>
          </div>
          <div>
            <h3 className="text-2xl font-bold">∞</h3>
            <p className="text-sm text-gray-500">Unlimited Access</p>
          </div>
          <div>
            <h3 className="text-2xl font-bold">Dedicated</h3>
            <p className="text-sm text-gray-500">CIR Connection</p>
          </div>
        </div>
      </div>

    </section>
  );
}
