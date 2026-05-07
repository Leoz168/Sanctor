"use client";

import Image from "next/image";
import { useState } from "react";
import { Check } from "lucide-react";
import { FormSection } from "@/components/create-post/form-section";

const memberCommunities = [
  {
    id: "computer-science",
    name: "Computer Science Collective",
    members: 1248,
    image: "/images/community-1.jpg",
  },
  {
    id: "student-market",
    name: "Student Market & Swap",
    members: 3450,
    image: "/images/community-5.jpg",
  },
  {
    id: "campus-social",
    name: "Campus Social Circle",
    members: 2100,
    image: "/images/community-4.jpg",
  },
  {
    id: "housing-leads",
    name: "Housing Leads Exchange",
    members: 876,
    image: "/images/listing-3.jpg",
  },
  {
    id: "roommate-search",
    name: "Roommate Search Hub",
    members: 1540,
    image: "/images/listing-4.jpg",
  },
];

export function CommunityShareSection() {
  const [selectedCommunities, setSelectedCommunities] = useState<string[]>([]);

  const toggleCommunity = (communityId: string) => {
    setSelectedCommunities((current) =>
      current.includes(communityId)
        ? current.filter((id) => id !== communityId)
        : [...current, communityId],
    );
  };

  return (
    <FormSection title="Share with Communities">
      <div>
        <p className="mb-8 text-lg font-semibold italic leading-8 text-gray-400">
          Broaden your reach by sharing this listing directly with the student
          circles you belong to.
        </p>

        <div className="max-h-[360px] overflow-y-auto pr-2">
          <div className="grid gap-4 lg:grid-cols-2">
            {memberCommunities.map((community) => {
              const isSelected = selectedCommunities.includes(community.id);

              return (
                <button
                  key={community.id}
                  type="button"
                  onClick={() => toggleCommunity(community.id)}
                  aria-pressed={isSelected}
                  className={`flex items-center gap-4 rounded-3xl border p-4 text-left transition-all ${
                    isSelected
                      ? "border-brand-orange bg-orange-50/60 shadow-lg shadow-brand-orange/10"
                      : "border-gray-100 bg-white hover:border-orange-200 hover:bg-brand-cream/60"
                  }`}
                >
                  <div className="relative h-16 w-16 shrink-0 overflow-hidden rounded-2xl bg-orange-50">
                    <Image
                      src={community.image}
                      alt={community.name}
                      fill
                      sizes="64px"
                      className="object-cover"
                    />
                  </div>

                  <div className="min-w-0 flex-1">
                    <h3 className="text-base font-black uppercase leading-tight text-gray-900">
                      {community.name}
                    </h3>
                    <p className="mt-1 text-sm font-bold italic text-gray-400">
                      {community.members.toLocaleString()} members
                    </p>
                  </div>

                  <span
                    className={`flex h-9 w-9 shrink-0 items-center justify-center rounded-xl border-2 transition-all ${
                      isSelected
                        ? "border-brand-orange bg-brand-orange text-white"
                        : "border-gray-100 bg-white text-transparent"
                    }`}
                  >
                    <Check className="h-5 w-5" />
                  </span>
                </button>
              );
            })}
          </div>
        </div>
      </div>
    </FormSection>
  );
}
