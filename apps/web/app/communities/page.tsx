"use client";

import Link from "next/link";
import { ChevronDown, Home, MessageSquare, Search, SlidersHorizontal, User } from "lucide-react";
import { CommunityCard } from "@/components/community-card";

const communities = [
  {
    id: 1,
    href: "/communities/computer-science-collective",
    name: "Computer Science Collective",
    description: "The primary hub for CSS/CS students to share resources, study groups, and tech news.",
    category: "ACADEMIC",
    members: 1248,
    postsPerWeek: 382,
    image: "/images/community-1.jpg",
    isJoined: true,
  },
  {
    id: 2,
    href: "/communities/residence-life-hub",
    name: "Residence Life Hub",
    description: "Stay updated on campus events, cafeteria menus, and dorm-specific announcements.",
    category: "RESIDENCE",
    members: 1892,
    postsPerWeek: 524,
    isJoined: false,
  },
  {
    id: 3,
    href: "/communities/student-market-swap",
    name: "Student Market & Swap",
    description: "Buy and sell textbooks, electronics, and furniture within the university community.",
    category: "MARKET",
    members: 3450,
    postsPerWeek: 890,
    isJoined: false,
  },
  {
    id: 4,
    href: "/communities/campus-social-circle",
    name: "Campus Social Circle",
    description: "Casual discussion for weekend plans, club events, and general campus life.",
    category: "SOCIAL",
    members: 2100,
    postsPerWeek: 412,
    image: "/images/community-4.jpg",
    isJoined: false,
  },
  {
    id: 5,
    href: "/communities/art-studio-zone",
    name: "The Art Studio Zone",
    description: "A creative space for visual artists, photographers, and designers to collaborate.",
    category: "SOCIAL",
    members: 850,
    postsPerWeek: 125,
    image: "/images/community-5.jpg",
    isJoined: false,
  },
  {
    id: 6,
    href: "/communities/pre-med-study-group",
    name: "Pre-Med Study Group",
    description: "Advanced study materials, MCAT prep, and lab report discussions.",
    category: "ACADEMIC",
    members: 1100,
    postsPerWeek: 245,
    isJoined: false,
  },
];

export default function CommunitiesPage() {
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
              <Link href="/communities" className="text-sm font-medium text-brand-orange transition-colors">
                Communities
              </Link>
              <Link href="/post-listings" className="text-sm font-medium text-gray-600 hover:text-brand-orange transition-colors">
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
                placeholder="Find communities..."
                className="w-full bg-brand-cream border border-gray-100 rounded-3xl pl-14 pr-6 py-4 text-base font-semibold focus:outline-none focus:bg-white focus:ring-2 focus:ring-brand-orange/20 transition-all shadow-inner"
              />
            </div>

            <div className="flex flex-col sm:flex-row gap-3">
              <label className="relative">
                <span className="sr-only">Community category</span>
                <select className="w-full sm:w-44 appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-5 py-4 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20">
                  <option>All Categories</option>
                  <option>Academic</option>
                  <option>Social</option>
                  <option>Residence</option>
                  <option>Market</option>
                </select>
                <ChevronDown className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" size={16} />
              </label>

              <label className="relative">
                <span className="sr-only">Sort communities</span>
                <select className="w-full sm:w-40 appearance-none bg-brand-cream border border-gray-100 rounded-2xl px-5 py-4 pr-10 text-sm font-bold text-gray-700 cursor-pointer hover:bg-white transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/20">
                  <option>Most Active</option>
                  <option>Newest</option>
                  <option>Largest</option>
                </select>
                <ChevronDown className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" size={16} />
              </label>

              <button className="h-[54px] w-full sm:w-[54px] rounded-2xl bg-brand-cream border border-gray-100 text-gray-700 hover:bg-white transition-colors flex items-center justify-center">
                <SlidersHorizontal size={20} />
                <span className="sr-only">More filters</span>
              </button>
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {communities.map((community) => (
              <CommunityCard key={community.id} {...community} />
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
        <MessageSquare className="w-5 h-5" />
        <span>Create a community</span>
      </button>
    </div>
  );
}
