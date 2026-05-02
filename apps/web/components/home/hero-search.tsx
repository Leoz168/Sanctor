"use client";

import { useState } from "react";
import { Search } from "lucide-react";
import { SelectControl } from "@/components/forms/select-control";

export function HeroSearch() {
  const [propertyType, setPropertyType] = useState("Property Type");

  return (
    <div className="relative max-w-5xl mx-auto" id="search-container">
      <div className="flex flex-col md:flex-row items-stretch gap-3 p-3 bg-white/75 backdrop-blur-md border border-white rounded-[2rem] shadow-2xl shadow-orange-900/10">
        <SelectControl
          label="Property type"
          value={propertyType}
          onChange={setPropertyType}
          options={["Property Type", "Apartment", "Shared House", "Studio", "Dorm"]}
          className="min-w-[220px]"
          variant="hero"
        />

        <div className="relative flex-grow">
          <div className="absolute left-6 top-1/2 -translate-y-1/2 text-gray-400">
            <Search size={28} />
          </div>
          <input
            type="text"
            placeholder="Search by school, campus, or city..."
            className="w-full bg-white border border-gray-100 rounded-2xl pl-16 pr-6 py-5 text-lg font-medium focus:outline-none focus:ring-2 focus:ring-brand-orange/20 transition-all shadow-md shadow-gray-900/5"
            id="search-input"
          />
        </div>

        <button
          className="bg-brand-orange text-white px-10 py-5 rounded-2xl text-lg font-bold flex items-center justify-center gap-3 hover:bg-orange-600 transition-all hover:shadow-xl hover:shadow-brand-orange/30 active:scale-95 md:min-w-[180px]"
          id="search-submit"
        >
          <Search size={26} />
          <span>Search</span>
        </button>
      </div>
    </div>
  );
}
