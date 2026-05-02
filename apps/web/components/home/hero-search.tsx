"use client";

import { useState } from "react";
import { Search } from "lucide-react";
import { SelectControl } from "@/components/forms/select-control";

export function HeroSearch() {
  const [propertyType, setPropertyType] = useState("Property Type");

  return (
    <div className="relative max-w-3xl mx-auto" id="search-container">
      <div className="flex flex-col md:flex-row items-stretch gap-2 p-2 bg-white/70 backdrop-blur-md border border-white rounded-3xl shadow-xl shadow-orange-900/5">
        <SelectControl
          label="Property type"
          value={propertyType}
          onChange={setPropertyType}
          options={["Property Type", "Apartment", "Shared House", "Studio", "Dorm"]}
          className="min-w-[160px]"
        />

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
  );
}
