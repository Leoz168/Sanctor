import Image from "next/image";
import Link from "next/link";
import {
  ArrowLeft,
  Bath,
  Bed,
  CalendarDays,
  CheckCircle2,
  Heart,
  Home,
  Mail,
  MapPin,
  MessageCircle,
  Share2,
  ShieldCheck,
  Sparkles,
  UserRound,
  Wifi,
} from "lucide-react";
import type { LucideIcon } from "lucide-react";
import { AppShell } from "@/components/layout/app-shell";
import { ListingPhotoCarousel } from "@/components/listing/listing-photo-carousel";

const listingImages = [
  { src: "/images/listing-1.jpg", alt: "Bright student studio bedroom" },
  { src: "/images/listing-2.jpg", alt: "Shared apartment kitchen" },
  { src: "/images/listing-3.jpg", alt: "Sunny apartment living area" },
  { src: "/images/listing-4.jpg", alt: "Quiet study and sleeping space" },
];

const listingStats = [
  { icon: Bed, label: "1 Bed" },
  { icon: Bath, label: "1 Bath" },
  { icon: Home, label: "Studio" },
  { icon: CalendarDays, label: "Fall term" },
];

const amenities = [
  "Furnished",
  "High-speed WiFi",
  "Laundry nearby",
  "Study desk",
  "Transit friendly",
  "Female friendly",
];

export default function CommunityListingPage() {
  return (
    <AppShell surface="cream">
      <div className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <Link
          href="/post-listings"
          className="mb-6 inline-flex items-center gap-2 text-sm font-bold text-gray-500 transition-colors hover:text-brand-orange"
        >
          <ArrowLeft size={18} />
          Back to listings
        </Link>

        <div className="grid items-stretch gap-8 lg:grid-cols-[1fr_380px]">
          <div className="flex h-full flex-col overflow-hidden rounded-[2rem] border border-gray-100 bg-white shadow-xl shadow-orange-900/5">
            <ListingPhotoCarousel images={listingImages} />

            <section className="min-h-[260px] border-t border-gray-100 p-7 sm:p-10">
              <h2 className="mb-4 text-xl font-bold text-gray-800">
                About this space
              </h2>
              <p className="text-base font-medium leading-8 text-gray-500">
                A bright, furnished studio close to campus with a quiet study
                setup, warm natural light, and quick access to transit, grocery
                stops, and student spaces. The listing is posted through a
                community thread so students can ask questions, verify context,
                and keep housing conversations connected.
              </p>
            </section>
          </div>

          <ListingInfoCard stats={listingStats} />
        </div>

        <div className="mt-8">
          <div className="space-y-8">
            <section className="rounded-[2rem] border border-gray-100 bg-white p-5 shadow-sm sm:p-6">
              <h2 className="mb-4 text-xl font-bold text-gray-800">
                Amenities
              </h2>
              <div className="flex flex-wrap gap-2.5">
                {amenities.map((amenity) => (
                  <div
                    key={amenity}
                    className="inline-flex items-center gap-2 rounded-full border border-orange-100 bg-orange-50/40 px-4 py-2.5"
                  >
                    <CheckCircle2 className="h-4 w-4 text-brand-orange" />
                    <span className="text-sm font-bold text-gray-700">{amenity}</span>
                  </div>
                ))}
              </div>
            </section>
          </div>
        </div>
      </div>
    </AppShell>
  );
}

interface ListingInfoCardProps {
  stats: typeof listingStats;
}

function ListingInfoCard({ stats }: ListingInfoCardProps) {
  return (
    <aside className="h-full">
      <div className="flex h-full flex-col rounded-[2rem] border border-gray-100 bg-white p-6 shadow-2xl shadow-orange-900/10">
        <div className="mb-5 flex items-start justify-between gap-4">
          <div className="min-w-0">
            <p className="mb-2 text-xs font-black uppercase tracking-[0.18em] text-brand-orange/70">
              Computer Science Collective
            </p>
            <h1 className="text-2xl font-bold leading-tight tracking-tight text-gray-900">
              Modern Studio near campus
            </h1>
          </div>
          <div className="flex shrink-0 gap-2">
            <button className="inline-flex h-10 w-10 items-center justify-center rounded-full border border-gray-100 bg-white text-gray-500 shadow-sm transition-colors hover:text-brand-orange">
              <Share2 size={17} />
            </button>
            <button className="inline-flex h-10 w-10 items-center justify-center rounded-full border border-gray-100 bg-white text-gray-500 shadow-sm transition-colors hover:text-brand-orange">
              <Heart size={17} />
            </button>
          </div>
        </div>

        <div className="mb-5 flex flex-col gap-2 text-sm font-bold text-gray-500">
          <span className="inline-flex items-center gap-2">
            <MapPin className="h-4 w-4 text-brand-orange" />
            St. George Campus, Toronto
          </span>
          <span className="inline-flex items-center gap-2">
            <ShieldCheck className="h-4 w-4 text-brand-orange" />
            Community verified
          </span>
        </div>

        <div className="mb-6 grid grid-cols-2 gap-2">
          {stats.map((stat) => (
            <div
              key={stat.label}
              className="flex items-center gap-2 rounded-2xl bg-brand-cream px-3 py-3"
            >
              <stat.icon className="h-4 w-4 text-brand-orange" />
              <span className="text-sm font-bold text-gray-700">
                {stat.label}
              </span>
            </div>
          ))}
        </div>

        <div className="mb-6 flex items-start justify-between border-t border-gray-100 pt-5">
          <div>
            <p className="text-xs font-black uppercase tracking-[0.18em] text-gray-400">
              Monthly rent
            </p>
            <div className="mt-2 flex items-baseline gap-1">
              <span className="text-4xl font-black text-brand-orange">$1200</span>
              <span className="font-bold text-gray-400">/mo</span>
            </div>
          </div>
          <span className="rounded-full bg-orange-50 px-3 py-1 text-xs font-black uppercase tracking-[0.12em] text-brand-orange">
            New
          </span>
        </div>

        <div className="space-y-4 border-y border-gray-100 py-5">
          <InfoRow icon={UserRound} label="Preferred roommate" value="Any gender" />
          <InfoRow icon={Wifi} label="Utilities" value="Included" />
          <InfoRow icon={Sparkles} label="Cleanliness" value="Weekly shared chores" />
          <InfoRow icon={Mail} label="Email" value="maya.k@school.edu" />
        </div>

        <div className="mt-auto py-5">
          <p className="mb-3 text-xs font-black uppercase tracking-[0.2em] text-gray-400">
            Posted by
          </p>
          <div className="flex items-center gap-3">
            <div className="relative h-12 w-12 overflow-hidden rounded-full bg-brand-cream">
              <Image
                src="/images/community-1.jpg"
                alt="Poster avatar"
                fill
                sizes="48px"
                className="object-cover"
              />
            </div>
            <div>
              <p className="font-black text-gray-900">Maya K.</p>
              <p className="text-sm font-semibold text-gray-400">
                Verified student poster
              </p>
            </div>
          </div>
        </div>

        <button className="inline-flex w-full items-center justify-center gap-2 rounded-2xl bg-brand-orange px-5 py-4 text-base font-black text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99]">
          <MessageCircle size={20} />
          Message Poster
        </button>
      </div>
    </aside>
  );
}

interface InfoRowProps {
  icon: LucideIcon;
  label: string;
  value: string;
}

function InfoRow({ icon: Icon, label, value }: InfoRowProps) {
  return (
    <div className="flex items-center justify-between gap-4">
      <div className="flex items-center gap-3 text-gray-500">
        <Icon className="h-5 w-5 text-brand-orange" />
        <span className="text-sm font-bold">{label}</span>
      </div>
      <span className="text-right text-sm font-black text-gray-800">{value}</span>
    </div>
  );
}
