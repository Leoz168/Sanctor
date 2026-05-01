import Image from "next/image";
import Link from "next/link";
import { MessageSquare } from "lucide-react";

interface CommunityCardProps {
  href: string;
  name: string;
  description: string;
  category: string;
  members: number;
  postsPerWeek: number;
  image?: string;
  isJoined?: boolean;
}

export function CommunityCard({
  href,
  name,
  description,
  category,
  members,
  postsPerWeek,
  image,
  isJoined = false,
}: CommunityCardProps) {
  return (
    <div className="bg-white rounded-[2rem] border border-gray-100 overflow-hidden group shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all">
      <div className="relative aspect-[16/10] bg-orange-50 overflow-hidden">
        {image ? (
          <Image
            src={image}
            alt={name}
            fill
            sizes="(min-width: 1024px) 33vw, (min-width: 768px) 50vw, 100vw"
            className="object-cover transition-transform duration-700 group-hover:scale-105"
          />
        ) : (
          <div className="h-full w-full flex items-center justify-center text-brand-orange">
            <MessageSquare size={40} />
          </div>
        )}
        <span className="absolute top-3 left-3 px-3 py-1 text-[10px] font-bold uppercase tracking-wider bg-gray-900/85 text-white rounded-full shadow-lg">
          {category}
        </span>
      </div>

      <div className="p-6">
        <Link href={href} className="block">
          <h3 className="text-lg font-bold text-gray-900 mb-2 group-hover:text-brand-orange transition-colors">
            {name}
          </h3>
        </Link>
        <p className="text-sm text-gray-500 mb-5 line-clamp-2 leading-relaxed">
          {description}
        </p>

        <div className="flex gap-8 mb-6">
          <div>
            <p className="text-lg font-bold text-gray-900">{members.toLocaleString()}</p>
            <p className="text-[10px] text-gray-400 uppercase tracking-widest font-bold">
              Members
            </p>
          </div>
          <div>
            <p className="text-lg font-bold text-gray-900">{postsPerWeek}</p>
            <p className="text-[10px] text-gray-400 uppercase tracking-widest font-bold">
              Posts/WK
            </p>
          </div>
        </div>

        <Link
          href={href}
          className={`block w-full rounded-xl px-4 py-3 text-center text-sm font-bold transition-colors ${
            isJoined
              ? "border border-gray-200 text-gray-500 hover:bg-gray-50"
              : "border border-brand-orange/30 bg-brand-orange/5 text-brand-orange hover:bg-brand-orange hover:text-white"
          }`}
        >
          {isJoined ? "JOINED" : "JOIN COMMUNITY"}
        </Link>
      </div>
    </div>
  );
}
