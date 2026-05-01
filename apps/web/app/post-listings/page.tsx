"use client";

import Link from "next/link";
import { ChevronDown, Filter, Home, Plus, Search, User } from "lucide-react";
import { ListingCard } from "@/components/listing-card";

const listings = [
  {
    id: 1,
    title: "Modern Studio near campus",
    price: 1200,
    location: "St. George Campus, Toronto",
    beds: 1,
    baths: 1,
    image: "/images/listing-1.jpg",
    badge: "featured" as const,
  },
  {
    id: 2,
    title: "Shared 3BR House - Female only",
    price: 850,
    location: "North Campus Area",
    beds: 3,
    baths: 2,
    image: "/images/listing-2.jpg",
    badge: "new" as const,
  },
  {
    id: 3,
    title: "Luxury Apartment in Downtown",
    price: 2100,
    location: "Downtown Core",
    beds: 2,
    baths: 2,
    image: "/images/listing-3.jpg",
  },
  {
    id: 4,
    title: "Cozy Loft for Students",
    price: 950,
    location: "East Side Campus",
    beds: 1,
    baths: 1,
    image: "/images/listing-4.jpg",
  },
  {
    id: 5,
    title: "Renovated Basement Suite",
    price: 1100,
    location: "West Campus Gardens",
    beds: 1,
    baths: 1,
    image: "/images/listing-5.jpg",
  },
  {
    id: 6,
    title: "Large 4BR Student Residence",
    price: 700,
    location: "Campus South",
    beds: 4,
    baths: 3,
    image: "/images/listing-6.jpg",
    badge: "new" as const,
  },
];

export default function PostListingsPage() {
  return (
    <div className="min-h-screen flex flex-col bg-white font-sans">
      <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-20 items-center">
            <Link href="/" className="flex items-center gap-2 group" id="nav-logo">
              <div className="p-2 bg-brand-orange rounded-xl text-white shadow-lg shadow-brand-orange/20">
                <Home size={24} />
              </div>
              <span className="text-2xl font-bold tracking-tight text-gray-900 group-hover:text-brand-orange transition-colors">
                Renting
              </span>
            </Link>

            <div className="hidden md:flex items-center gap-8">
              <a href="#" className="text-sm font-medium text-gray-600 hover:text-brand-orange transition-colors">
                Communities
              </a>
              <Link href="/post-listings" className="text-sm font-medium text-brand-orange transition-colors">
                Post Listing
              </Link>
              <div className="h-4 w-px bg-gray-200" />
              <button className="flex items-center gap-2 px-5 py-2.5 bg-brand-orange text-white rounded-full font-medium shadow-md hover:shadow-lg hover:bg-orange-600 transition-all active:scale-95">
                <User size={18} />
                <span>Login</span>
              </button>
            </div>
          </div>
        </div>
      </nav>

      <main className="flex-1">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
          <div className="flex flex-col lg:flex-row lg:items-center gap-4 mb-10">
            <div className="relative flex-1">
              <Search className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
              <input
                type="text"
                placeholder="Search postings..."
                className="w-full bg-brand-cream border border-gray-100 rounded-3xl pl-14 pr-6 py-4 text-base font-semibold focus:outline-none focus:bg-white focus:ring-2 focus:ring-brand-orange/20 transition-all shadow-inner"
              />
            </div>

            <div className="flex flex-col sm:flex-row gap-3">
              <label className="relative">
                <span className="sr-only">Listing type</span>
                <select className="w-full sm:w-40 appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-5 py-4 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20">
                  <option>All Types</option>
                  <option>Apartment</option>
                  <option>House</option>
                  <option>Studio</option>
                  <option>Shared</option>
                </select>
                <ChevronDown className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" size={16} />
              </label>

              <label className="relative">
                <span className="sr-only">Price range</span>
                <select className="w-full sm:w-40 appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-5 py-4 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20">
                  <option>Any Price</option>
                  <option>Under $500</option>
                  <option>Under $1000</option>
                  <option>Under $1500</option>
                  <option>Under $2000</option>
                </select>
                <ChevronDown className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" size={16} />
              </label>

              <button className="h-[54px] w-full sm:w-[54px] rounded-2xl bg-brand-cream border border-gray-100 text-gray-700 hover:bg-white transition-colors flex items-center justify-center">
                <Filter size={20} />
                <span className="sr-only">More filters</span>
              </button>
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {listings.map((listing) => (
              <ListingCard key={listing.id} {...listing} />
            ))}
          </div>
        </div>
      </main>

      <footer className="bg-gray-50 border-t border-gray-200 pt-16 pb-8">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row justify-between items-center gap-8 mb-12">
            <Link href="/" className="flex items-center gap-2">
              <div className="p-1.5 bg-brand-orange rounded-lg text-white">
                <Home size={18} />
              </div>
              <span className="text-xl font-bold tracking-tight">Renting</span>
            </Link>
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

      <button className="fixed bottom-8 right-6 sm:right-8 flex items-center gap-2 px-5 py-4 bg-brand-orange text-white rounded-full font-bold shadow-xl shadow-brand-orange/30 hover:bg-orange-600 transition-all active:scale-95">
        <Plus className="w-5 h-5" />
        <span>Make a post</span>
      </button>
    </div>
  );
}
