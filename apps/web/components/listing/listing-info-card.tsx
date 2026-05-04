import Image from "next/image";
import {
  Heart,
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

interface ListingStat {
  icon: LucideIcon;
  label: string;
}

interface ListingInfoRow {
  icon: LucideIcon;
  label: string;
  value: string;
}

interface Poster {
  name: string;
  status: string;
  avatar: string;
  email: string;
}

interface ListingInfoCardProps {
  communityName: string;
  title: string;
  location: string;
  verifiedLabel: string;
  monthlyRent: number;
  badge: string;
  stats: ListingStat[];
  poster: Poster;
}

export function ListingInfoCard({
  communityName,
  title,
  location,
  verifiedLabel,
  monthlyRent,
  badge,
  stats,
  poster,
}: ListingInfoCardProps) {
  const infoRows: ListingInfoRow[] = [
    { icon: UserRound, label: "Preferred roommate", value: "Any gender" },
    { icon: Wifi, label: "Utilities", value: "Included" },
    { icon: Sparkles, label: "Cleanliness", value: "Weekly shared chores" },
    { icon: Mail, label: "Email", value: poster.email },
  ];

  return (
    <aside className="h-full">
      <div className="flex h-full flex-col rounded-[2rem] border border-gray-100 bg-white p-6 shadow-2xl shadow-orange-900/10">
        <div className="mb-5 flex items-start justify-between gap-4">
          <div className="min-w-0">
            <p className="mb-2 text-xs font-black uppercase tracking-[0.18em] text-brand-orange/70">
              {communityName}
            </p>
            <h1 className="text-2xl font-bold leading-tight tracking-tight text-gray-900">
              {title}
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
            {location}
          </span>
          <span className="inline-flex items-center gap-2">
            <ShieldCheck className="h-4 w-4 text-brand-orange" />
            {verifiedLabel}
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
              <span className="text-4xl font-black text-brand-orange">
                ${monthlyRent}
              </span>
              <span className="font-bold text-gray-400">/mo</span>
            </div>
          </div>
          <span className="rounded-full bg-orange-50 px-3 py-1 text-xs font-black uppercase tracking-[0.12em] text-brand-orange">
            {badge}
          </span>
        </div>

        <div className="space-y-4 border-y border-gray-100 py-5">
          {infoRows.map((row) => (
            <InfoRow key={row.label} {...row} />
          ))}
        </div>

        <div className="mt-auto py-5">
          <p className="mb-3 text-xs font-black uppercase tracking-[0.2em] text-gray-400">
            Posted by
          </p>
          <div className="flex items-center gap-3">
            <div className="relative h-12 w-12 overflow-hidden rounded-full bg-brand-cream">
              <Image
                src={poster.avatar}
                alt="Poster avatar"
                fill
                sizes="48px"
                className="object-cover"
              />
            </div>
            <div>
              <p className="font-black text-gray-900">{poster.name}</p>
              <p className="text-sm font-semibold text-gray-400">
                {poster.status}
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

function InfoRow({ icon: Icon, label, value }: ListingInfoRow) {
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
