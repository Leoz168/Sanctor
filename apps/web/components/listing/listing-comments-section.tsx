"use client";

import Image from "next/image";
import { useState } from "react";
import { ChevronDown, MessageCircle, Send } from "lucide-react";

const comments = [
  {
    id: 1,
    name: "Nadia Chen",
    role: "Verified community member",
    time: "2 hours ago",
    avatar: "/images/community-4.jpg",
    text: "I toured this building last term. The study lounge is quiet and the walk to St. George is very manageable.",
  },
  {
    id: 2,
    name: "Aarav Patel",
    role: "Roommate seeker",
    time: "Yesterday",
    avatar: "/images/community-5.jpg",
    text: "Does anyone know if utilities are actually included year-round, or only during the fall sublet period?",
  },
  {
    id: 3,
    name: "Maya K.",
    role: "Poster",
    time: "Yesterday",
    avatar: "/images/community-1.jpg",
    text: "Utilities are included for the full lease period. Happy to share the building rules with anyone interested.",
  },
];

export function ListingCommentsSection() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <section className="overflow-hidden rounded-[2rem] border border-gray-100 bg-white shadow-sm">
      <button
        type="button"
        onClick={() => setIsOpen((open) => !open)}
        className="flex w-full items-center justify-between gap-4 px-5 py-5 text-left transition-colors hover:bg-brand-cream/60 sm:px-6"
        aria-expanded={isOpen}
      >
        <div className="flex items-center gap-4">
          <div className="flex h-12 w-12 items-center justify-center rounded-2xl bg-orange-50 text-brand-orange">
            <MessageCircle className="h-5 w-5" />
          </div>
          <div>
            <h2 className="text-xl font-bold text-gray-800">Comments</h2>
            <p className="mt-1 text-sm font-semibold text-gray-400">
              {comments.length} community replies about this listing
            </p>
          </div>
        </div>

        <ChevronDown
          className={`h-5 w-5 shrink-0 text-gray-400 transition-transform ${
            isOpen ? "rotate-180" : ""
          }`}
        />
      </button>

      <div
        className={`grid transition-[grid-template-rows] duration-300 ease-out ${
          isOpen ? "grid-rows-[1fr]" : "grid-rows-[0fr]"
        }`}
      >
        <div className="overflow-hidden">
          <div className="border-t border-gray-100 px-5 py-5 sm:px-6">
            <div className="space-y-4">
              {comments.map((comment) => (
                <article
                  key={comment.id}
                  className="rounded-3xl border border-orange-100/70 bg-brand-cream/45 p-4"
                >
                  <div className="flex gap-3">
                    <div className="relative h-11 w-11 shrink-0 overflow-hidden rounded-2xl bg-orange-50">
                      <Image
                        src={comment.avatar}
                        alt={comment.name}
                        fill
                        sizes="44px"
                        className="object-cover"
                      />
                    </div>
                    <div className="min-w-0 flex-1">
                      <div className="flex flex-wrap items-baseline gap-x-2 gap-y-1">
                        <h3 className="font-bold text-gray-900">
                          {comment.name}
                        </h3>
                        <span className="text-xs font-bold uppercase tracking-[0.16em] text-brand-orange">
                          {comment.role}
                        </span>
                        <span className="text-xs font-semibold text-gray-400">
                          {comment.time}
                        </span>
                      </div>
                      <p className="mt-2 text-sm font-medium leading-6 text-gray-600">
                        {comment.text}
                      </p>
                    </div>
                  </div>
                </article>
              ))}
            </div>

            <form className="mt-5 flex flex-col gap-3 rounded-3xl border border-gray-100 bg-white p-3 shadow-inner shadow-gray-900/5 sm:flex-row sm:items-center">
              <input
                type="text"
                placeholder="Ask a question or add helpful context..."
                className="min-h-12 flex-1 rounded-2xl bg-gray-50 px-4 text-sm font-semibold text-gray-800 outline-none placeholder:text-gray-400 focus:bg-brand-cream"
              />
              <button
                type="button"
                className="inline-flex items-center justify-center gap-2 rounded-2xl bg-brand-orange px-5 py-3 text-sm font-bold text-white shadow-lg shadow-brand-orange/20 transition-all hover:bg-orange-600 active:scale-95"
              >
                <Send className="h-4 w-4" />
                Comment
              </button>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
}
