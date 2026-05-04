import Image from "next/image";
import { MessageSquare, ShieldCheck, Users } from "lucide-react";
import type { LucideIcon } from "lucide-react";

interface CommunityDetailHeroProps {
  name: string;
  category: string;
  description: string;
  image: string;
  members: number;
  postsPerWeek: number;
}

export function CommunityDetailHero({
  name,
  category,
  description,
  image,
  members,
  postsPerWeek,
}: CommunityDetailHeroProps) {
  return (
    <section className="overflow-hidden rounded-[2rem] border border-gray-100 bg-white shadow-xl shadow-orange-900/5">
      <div className="grid lg:grid-cols-[0.92fr_1.08fr]">
        <div className="flex flex-col justify-center p-7 sm:p-10">
          <p className="mb-4 text-xs font-black uppercase tracking-[0.22em] text-brand-orange/70">
            {category}
          </p>
          <h1 className="max-w-2xl text-4xl font-black tracking-tight text-gray-950 sm:text-5xl">
            {name}
          </h1>
          <p className="mt-5 max-w-2xl text-base font-medium leading-8 text-gray-500">
            {description}
          </p>

          <div className="mt-8 grid gap-3 sm:grid-cols-3">
            <HeroStat icon={Users} label="Members" value={members.toLocaleString()} />
            <HeroStat icon={MessageSquare} label="Posts / week" value={postsPerWeek.toString()} />
            <HeroStat icon={ShieldCheck} label="Status" value="Verified" />
          </div>

          <button className="mt-8 inline-flex w-fit items-center justify-center rounded-full bg-brand-orange px-7 py-3.5 text-sm font-black text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99]">
            Join Community
          </button>
        </div>

        <div className="relative min-h-[320px] lg:min-h-[520px]">
          <Image
            src={image}
            alt={name}
            fill
            priority
            sizes="(min-width: 1024px) 50vw, 100vw"
            className="object-cover"
          />
        </div>
      </div>
    </section>
  );
}

interface HeroStatProps {
  icon: LucideIcon;
  label: string;
  value: string;
}

function HeroStat({ icon: Icon, label, value }: HeroStatProps) {
  return (
    <div className="rounded-2xl bg-brand-cream px-4 py-4">
      <Icon className="mb-3 h-5 w-5 text-brand-orange" />
      <p className="text-lg font-black text-gray-900">{value}</p>
      <p className="text-[10px] font-black uppercase tracking-[0.18em] text-gray-400">
        {label}
      </p>
    </div>
  );
}
