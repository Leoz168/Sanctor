import Link from "next/link";
import { ArrowLeft } from "lucide-react";
import { CreatePostForm } from "@/components/create-post/create-post-form";
import { BrandLogo } from "@/components/navigation/brand-logo";

export default function CreatePostPage() {
  return (
    <main className="min-h-screen bg-white px-4 py-8 font-sans text-[#1A1A1A]">
      <div className="mx-auto max-w-5xl">
        <div className="mb-8 flex items-center justify-between gap-4">
          <BrandLogo />

          <Link
            href="/post-listings"
            className="inline-flex items-center gap-2 text-sm font-bold text-gray-400 hover:text-brand-orange"
          >
            <ArrowLeft size={18} />
            To listings
          </Link>
        </div>

        <header className="mb-12">
          <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
            Create a <span className="text-brand-orange">Listing</span>
          </h1>
          <p className="mt-3 text-lg font-medium italic text-gray-500">
            Fill in the details below to find your perfect tenant or roommate.
          </p>
        </header>

        <CreatePostForm />
      </div>
    </main>
  );
}
