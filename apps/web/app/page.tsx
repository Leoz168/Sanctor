"use client";

import { useState } from "react";
import {
  ChevronDown,
  Home,
  Mail,
  Menu,
  MessageSquare,
  Search,
  ShoppingBag,
  User,
  X,
} from "lucide-react";

export default function HomePage() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [propertyType, setPropertyType] = useState("Property Type");

  return (
    <div className="min-h-screen flex flex-col font-sans">
      <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-20 items-center">
            <div className="flex items-center gap-2 group cursor-pointer" id="nav-logo">
              <div className="p-2 bg-brand-orange rounded-xl text-white shadow-lg shadow-brand-orange/20">
                <Home size={24} />
              </div>
              <span className="text-2xl font-bold tracking-tight text-gray-900 group-hover:text-brand-orange transition-colors">
                Renting
              </span>
            </div>

            <div className="hidden md:flex items-center gap-8">
              <a
                href="#"
                className="text-sm font-medium text-gray-600 hover:text-brand-orange transition-colors"
                id="nav-communities"
              >
                Communities
              </a>
              <a
                href="#"
                className="text-sm font-medium text-gray-600 hover:text-brand-orange transition-colors"
                id="nav-post-listing"
              >
                Post Listing
              </a>
              <div className="h-4 w-px bg-gray-200" />
              <button
                className="flex items-center gap-2 px-5 py-2.5 bg-brand-orange text-white rounded-full font-medium shadow-md hover:shadow-lg hover:bg-orange-600 transition-all active:scale-95"
                id="nav-login"
              >
                <User size={18} />
                <span>Login</span>
              </button>
            </div>

            <div className="md:hidden">
              <button
                onClick={() => setIsMenuOpen(!isMenuOpen)}
                className="p-2 text-gray-600 hover:text-brand-orange transition-colors"
                id="mobile-menu-toggle"
                aria-label="Toggle menu"
              >
                {isMenuOpen ? <X size={24} /> : <Menu size={24} />}
              </button>
            </div>
          </div>
        </div>

        {isMenuOpen && (
          <div className="md:hidden bg-white border-b border-gray-200 p-4 space-y-4">
            <a href="#" className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-brand-orange">
              Communities
            </a>
            <a href="#" className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-brand-orange">
              Post Listing
            </a>
            <button className="w-full flex items-center justify-center gap-2 px-5 py-3 bg-brand-orange text-white rounded-full font-medium">
              <User size={18} />
              <span>Login</span>
            </button>
          </div>
        )}
      </nav>

      <main className="flex-grow">
        <section className="relative pt-20 pb-32 overflow-hidden px-4">
          <div className="absolute top-0 left-1/2 -translate-x-1/2 w-full max-w-7xl h-full -z-10 bg-[radial-gradient(circle_at_50%_0%,rgba(255,237,213,0.5),transparent_65%)] pointer-events-none" />

          <div className="max-w-4xl mx-auto text-center" id="hero-content">
            <div>
              <h1 className="text-5xl md:text-7xl font-bold tracking-tight text-gray-900 mb-6 leading-[1.1]">
                Welcome to <span className="text-brand-orange">Renting!!!</span>
              </h1>
              <p className="text-xl md:text-2xl text-gray-600 mb-12 max-w-2xl mx-auto font-medium">
                Find the right housing for you! Explore campus communities and trusted listings.
              </p>
            </div>

            <div className="relative max-w-3xl mx-auto" id="search-container">
              <div className="flex flex-col md:flex-row items-stretch gap-2 p-2 bg-white/70 backdrop-blur-md border border-white rounded-3xl shadow-xl shadow-orange-900/5">
                <div className="relative min-w-[160px]">
                  <select
                    value={propertyType}
                    onChange={(event) => setPropertyType(event.target.value)}
                    className="w-full appearance-none bg-white/50 border border-gray-100 rounded-2xl px-4 py-4 text-sm font-semibold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20"
                    id="property-type-select"
                  >
                    <option>Property Type</option>
                    <option>Apartment</option>
                    <option>Shared House</option>
                    <option>Studio</option>
                    <option>Dorm</option>
                  </select>
                  <ChevronDown
                    className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none"
                    size={16}
                  />
                </div>

                <div className="relative flex-grow">
                  <div className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400">
                    <Search size={20} />
                  </div>
                  <input
                    type="text"
                    placeholder="Search by school, campus, or city..."
                    className="w-full bg-white border border-gray-100 rounded-2xl pl-12 pr-4 py-4 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-brand-orange/20 transition-all shadow-sm"
                    id="search-input"
                  />
                </div>

                <button
                  className="bg-brand-orange text-white px-8 py-4 rounded-2xl font-bold flex items-center justify-center gap-2 hover:bg-orange-600 transition-all hover:shadow-lg hover:shadow-brand-orange/30 active:scale-95"
                  id="search-submit"
                >
                  <Search size={20} />
                  <span>Search</span>
                </button>
              </div>
            </div>
          </div>
        </section>

        <section className="max-w-7xl mx-auto px-4 pb-24">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8" id="feature-columns">
            <div
              className="bg-white p-8 rounded-[2rem] border border-gray-100 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all group"
              id="feature-subscribe"
            >
              <div className="w-14 h-14 bg-orange-50 text-brand-orange rounded-2xl flex items-center justify-center mb-6 group-hover:bg-brand-orange group-hover:text-white transition-colors">
                <Mail size={28} />
              </div>
              <h3 className="text-2xl font-bold mb-4">Subscribe</h3>
              <p className="text-gray-600 leading-relaxed mb-6">
                Get notified as soon as new property listings that match your criteria become available.
              </p>
              <div className="relative">
                <input
                  type="email"
                  placeholder="your@email.com"
                  className="w-full bg-gray-50 border border-gray-100 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-brand-orange/20 mb-3"
                />
                <button className="w-full py-3 bg-gray-900 text-white rounded-xl text-sm font-bold hover:bg-black transition-colors">
                  Notify Me
                </button>
              </div>
            </div>

            <div
              className="bg-white p-8 rounded-[2rem] border border-gray-100 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all group"
              id="feature-forums"
            >
              <div className="w-14 h-14 bg-orange-50 text-brand-orange rounded-2xl flex items-center justify-center mb-6 group-hover:bg-brand-orange group-hover:text-white transition-colors">
                <MessageSquare size={28} />
              </div>
              <h3 className="text-2xl font-bold mb-4">Forums</h3>
              <p className="text-gray-600 leading-relaxed mb-6">
                Connect with other students, find roommates, and discuss neighborhood safety or landlords.
              </p>
              <ul className="space-y-3">
                {["Roommate Search", "General Housing Discussion", "Landlord Reviews"].map((topic) => (
                  <li
                    key={topic}
                    className="flex items-center gap-2 text-sm font-medium text-gray-500 hover:text-brand-orange cursor-pointer"
                  >
                    <div className="w-1.5 h-1.5 rounded-full bg-brand-orange/30" />
                    {topic}
                  </li>
                ))}
              </ul>
            </div>

            <div
              className="bg-white p-8 rounded-[2rem] border border-gray-100 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all group"
              id="feature-market"
            >
              <div className="w-14 h-14 bg-orange-50 text-brand-orange rounded-2xl flex items-center justify-center mb-6 group-hover:bg-brand-orange group-hover:text-white transition-colors">
                <ShoppingBag size={28} />
              </div>
              <h3 className="text-2xl font-bold mb-4">Market</h3>
              <p className="text-gray-600 leading-relaxed mb-6">
                Buy and sell second-hand furniture, textbooks, and essentials specifically for student living.
              </p>
              <div className="grid grid-cols-2 gap-3">
                <div className="h-20 bg-gray-50 rounded-xl flex items-center justify-center border border-dashed border-gray-200">
                  <span className="text-[10px] uppercase tracking-widest font-bold text-gray-400">Furniture</span>
                </div>
                <div className="h-20 bg-gray-50 rounded-xl flex items-center justify-center border border-dashed border-gray-200">
                  <span className="text-[10px] uppercase tracking-widest font-bold text-gray-400">Books</span>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="bg-gray-50 border-t border-gray-200 pt-16 pb-8">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row justify-between items-center gap-8 mb-12">
            <div className="flex items-center gap-2">
              <div className="p-1.5 bg-brand-orange rounded-lg text-white">
                <Home size={18} />
              </div>
              <span className="text-xl font-bold tracking-tight">Renting</span>
            </div>
            <div className="flex gap-8 text-sm font-medium text-gray-500">
              <a href="#" className="hover:text-brand-orange">
                Terms
              </a>
              <a href="#" className="hover:text-brand-orange">
                Privacy
              </a>
              <a href="#" className="hover:text-brand-orange">
                Contact
              </a>
              <a href="#" className="hover:text-brand-orange">
                Help
              </a>
            </div>
          </div>
          <div className="text-center text-gray-400 text-xs font-medium">
            &copy; {new Date().getFullYear()} Renting. Dedicated to student housing solutions.
          </div>
        </div>
      </footer>
    </div>
  );
}
