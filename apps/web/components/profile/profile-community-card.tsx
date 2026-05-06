import Image from "next/image";
import Link from "next/link";
import { ArrowRight, Star, Users } from "lucide-react";
import { ProfileCommunityNotificationButton } from "@/components/profile/profile-community-notification-button";

interface ProfileCommunityCardProps {
  name: string;
  description: string;
  image: string;
  href: string;
  members: number;
  role: string;
  isLead?: boolean;
}

export function ProfileCommunityCard({
  name,
  description,
  image,
  href,
  members,
  role,
  isLead = false,
}: ProfileCommunityCardProps) {
  return (
    <article className="mx-auto flex h-full max-w-[350px] flex-col overflow-hidden rounded-[1.75rem] border border-orange-100 bg-white shadow-sm shadow-orange-900/5 transition-all hover:-translate-y-1 hover:shadow-xl hover:shadow-orange-900/10">
      <div className="relative aspect-[16/9] overflow-hidden bg-orange-50">
        <Image
          src={image}
          alt={name}
          fill
          sizes="(min-width: 1024px) 330px, 100vw"
          className="object-cover"
        />
        <div className="absolute inset-x-0 bottom-0 h-20 bg-gradient-to-t from-gray-900/55 to-transparent" />

        <span
          className={`absolute left-3 top-3 inline-flex items-center gap-1.5 rounded-full px-3 py-1.5 text-[9px] font-bold uppercase tracking-[0.14em] shadow-md ${
            isLead
              ? "bg-gray-900 text-white"
              : "bg-white text-brand-orange"
          }`}
        >
          {isLead ? <Star className="h-3 w-3 fill-current" /> : null}
          {role}
        </span>

        <span className="absolute bottom-3 left-3 inline-flex items-center gap-1.5 rounded-full bg-white/85 px-3 py-1.5 text-[9px] font-bold uppercase tracking-[0.12em] text-gray-700 shadow-md backdrop-blur-sm">
          <Users className="h-3 w-3 text-brand-orange" />
          {members.toLocaleString()} members
        </span>
      </div>

      <div className="flex flex-1 flex-col p-5">
        <div className="flex items-start justify-between gap-3">
          <h2 className="max-w-[13rem] text-[1.35rem] font-bold leading-tight tracking-tight text-gray-900">
            {name}
          </h2>
          <ProfileCommunityNotificationButton />
        </div>

        <p className="mt-3 text-[13px] font-semibold italic leading-5 text-gray-500">
          "{description}"
        </p>

        <div className="mt-auto flex items-center justify-between border-t border-orange-100/70 pt-5">
          <span className="inline-flex items-center gap-2 text-[9px] font-bold uppercase tracking-[0.12em] text-brand-orange">
            <span className="h-2 w-2 rounded-full bg-brand-orange" />
            Active now
          </span>

          <Link
            href={href}
            className="inline-flex items-center gap-2 text-[10px] font-bold uppercase tracking-[0.14em] text-gray-900 transition-colors hover:text-brand-orange"
          >
            Enter
            <ArrowRight className="h-4 w-4" />
          </Link>
        </div>
      </div>
    </article>
  );
}
