import Image from "next/image";
import Link from "next/link";
import { MapPin } from "lucide-react";

export function ProfileListingManagementCard() {
  return (
    <article className="rounded-[2rem] border border-orange-100 bg-white/95 p-6 shadow-sm shadow-orange-900/5">
      <div className="grid gap-6 lg:grid-cols-[300px_1fr] lg:items-center">
        <div className="relative aspect-[4/3] overflow-hidden rounded-[1.5rem] bg-orange-50">
          <Image
            src="/images/listing-1.jpg"
            alt="Modern Studio near campus"
            fill
            sizes="300px"
            className="object-cover"
          />
          <span className="absolute left-4 top-4 rounded-full bg-white px-4 py-2 text-xs font-bold uppercase tracking-[0.16em] text-brand-orange shadow-md">
            Active
          </span>
        </div>

        <div className="min-w-0">
          <div className="flex flex-col gap-4 xl:flex-row xl:items-start xl:justify-between">
            <div>
              <h2 className="text-2xl font-bold tracking-tight text-gray-900">
                Modern Studio near campus
              </h2>
              <p className="mt-3 flex items-center gap-2 text-base font-semibold italic text-gray-500">
                <MapPin className="h-5 w-5 text-brand-orange" />
                St. George Campus, Toronto
              </p>
            </div>

            <p className="shrink-0 text-3xl font-bold text-brand-orange">
              $1200<span className="text-base text-gray-400">/mo</span>
            </p>
          </div>

          <div className="mt-8 flex flex-col gap-3 sm:flex-row sm:items-center">
            <Link
              href="/communities/computer-science-collective/listings/modern-studio-near-campus"
              className="inline-flex h-12 items-center justify-center rounded-2xl bg-gray-900 px-7 text-xs font-bold uppercase tracking-[0.16em] text-white shadow-xl shadow-gray-900/15 transition-all hover:bg-brand-orange hover:shadow-brand-orange/20"
            >
              Manage details
            </Link>
            <button
              type="button"
              className="inline-flex h-12 items-center justify-center rounded-2xl border border-orange-100 bg-white px-7 text-xs font-bold uppercase tracking-[0.16em] text-gray-400 transition-all hover:border-brand-orange hover:text-brand-orange"
            >
              Archive post
            </button>
          </div>

          <p className="mt-6 text-right text-xs font-bold uppercase tracking-[0.18em] text-gray-300">
            Added 2026-05-03
          </p>
        </div>
      </div>
    </article>
  );
}
