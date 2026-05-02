import Link from "next/link";
import Image from "next/image";
import { Plus } from "lucide-react";

export function PostListingCta() {
  return (
    <section className="mx-auto max-w-6xl px-4 pb-20">
      <div className="grid overflow-hidden rounded-[2rem] bg-brand-orange shadow-2xl shadow-orange-900/10 lg:grid-cols-2">
        <div className="flex min-h-[320px] flex-col items-center justify-center px-8 py-12 text-center text-white sm:px-12">
          <h2 className="text-4xl font-black leading-none tracking-tight sm:text-5xl">
            Own a space?
            <span className="mt-2 block text-white/70 italic">List it.</span>
          </h2>
          <p className="mx-auto mt-6 max-w-sm text-base font-semibold italic leading-relaxed text-white/90">
            Join the community of students finding reliable roommates every day. It&apos;s free, fast, and secure.
          </p>
          <Link
            href="/create-post"
            className="mx-auto mt-8 inline-flex items-center gap-3 rounded-2xl bg-white px-8 py-4 text-sm font-black uppercase tracking-[0.18em] text-brand-orange shadow-xl shadow-orange-900/10 transition-all hover:bg-orange-50 active:scale-[0.99]"
          >
            Post a listing
            <span className="flex h-7 w-7 items-center justify-center rounded-full bg-brand-orange text-white">
              <Plus size={18} />
            </span>
          </Link>
        </div>

        <div className="relative min-h-[260px] lg:min-h-[360px]">
          <Image
            src="/images/listing-3.jpg"
            alt="Modern student rental property"
            fill
            sizes="(min-width: 1024px) 50vw, 100vw"
            className="object-cover"
          />
        </div>
      </div>
    </section>
  );
}
