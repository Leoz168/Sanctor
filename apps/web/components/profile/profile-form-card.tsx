import Image from "next/image";
import { Camera, GraduationCap, Home, Mail, Plus, ShieldCheck } from "lucide-react";

const profileFields = [
  { label: "First Name", value: "Alex" },
  { label: "Last Name", value: "Rivera" },
  { label: "Gender", value: "Male" },
  { label: "Age", value: "21", type: "number" },
  { label: "Institution", value: "University of Toronto" },
  { label: "Major", value: "Architecture & Design" },
];

const accountHighlights = [
  { label: "Verified student", icon: ShieldCheck },
  { label: "4 communities", icon: GraduationCap },
  { label: "3 active listings", icon: Home },
];

export function ProfileFormCard() {
  return (
    <section className="overflow-hidden rounded-[2rem] border border-orange-100/70 bg-white shadow-2xl shadow-orange-900/8">
      <div className="h-24 bg-[linear-gradient(135deg,#fff4e6_0%,#fffaf2_45%,#f27d26_100%)]" />

      <div className="px-6 pb-8 sm:px-8 lg:px-12 lg:pb-12">
        <div className="-mt-12 flex flex-col gap-6 md:flex-row md:items-end md:justify-between">
          <div className="flex flex-col gap-5 sm:flex-row sm:items-end">
            <div className="relative h-36 w-36 shrink-0 rounded-[2rem] border-8 border-orange-50 bg-orange-50 shadow-xl shadow-orange-900/10">
              <Image
                src="/images/community-4.jpg"
                alt="Alex Rivera profile photo"
                fill
                sizes="144px"
                className="rounded-[1.5rem] object-cover"
              />
              <button
                type="button"
                className="absolute -bottom-3 -right-3 flex h-12 w-12 items-center justify-center rounded-2xl border-4 border-white bg-gray-900 text-white shadow-lg transition-transform hover:scale-105"
                aria-label="Upload profile photo"
              >
                <Plus className="h-6 w-6" />
              </button>
            </div>

            <div className="pb-2">
              <p className="mb-2 inline-flex items-center gap-2 rounded-full bg-orange-50 px-4 py-2 text-xs font-bold uppercase tracking-[0.18em] text-brand-orange">
                <Camera className="h-4 w-4" />
                Student profile
              </p>
              <h1 className="text-4xl font-bold tracking-tight text-gray-900">
                Alex Rivera
              </h1>
              <p className="mt-2 text-lg font-semibold italic text-gray-400">
                Student at University of Toronto
              </p>
            </div>
          </div>

          <div className="flex flex-wrap gap-2">
            {accountHighlights.map((highlight) => {
              const Icon = highlight.icon;

              return (
                <span
                  key={highlight.label}
                  className="inline-flex items-center gap-2 rounded-full border border-orange-100 bg-brand-cream px-4 py-2 text-sm font-bold text-gray-600"
                >
                  <Icon className="h-4 w-4 text-brand-orange" />
                  {highlight.label}
                </span>
              );
            })}
          </div>
        </div>

        <div className="mt-10 grid gap-6 md:grid-cols-2">
          {profileFields.map((field) => (
            <label key={field.label} className="block">
              <span className="mb-2 block text-[11px] font-bold uppercase tracking-[0.24em] text-gray-400">
                {field.label}
              </span>
              <input
                type={field.type ?? "text"}
                defaultValue={field.value}
                className="h-16 w-full rounded-2xl border border-gray-100 bg-gray-50 px-6 text-base font-bold text-gray-800 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10"
              />
            </label>
          ))}
        </div>

        <div className="mt-6 grid gap-6 lg:grid-cols-[1.2fr_0.8fr]">
          <label className="block">
            <span className="mb-2 block text-[11px] font-bold uppercase tracking-[0.24em] text-gray-400">
              Bio
            </span>
            <textarea
              defaultValue="Looking for clean, transit-friendly housing near campus. I like quiet study spaces, shared dinners, and communities that keep things organized."
              rows={5}
              className="w-full resize-none rounded-3xl border border-gray-100 bg-gray-50 px-6 py-5 text-base font-semibold leading-relaxed text-gray-700 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10"
            />
          </label>

          <div className="rounded-3xl border border-orange-100 bg-brand-cream p-6">
            <div className="flex h-12 w-12 items-center justify-center rounded-2xl bg-white text-brand-orange shadow-sm">
              <Mail className="h-5 w-5" />
            </div>
            <p className="mt-5 text-xs font-bold uppercase tracking-[0.22em] text-gray-400">
              Contact email
            </p>
            <p className="mt-2 text-lg font-bold text-gray-900">
              alex.rivera@school.edu
            </p>
            <p className="mt-3 text-sm font-medium leading-relaxed text-gray-500">
              This email is only shown to verified renters and community members.
            </p>
          </div>
        </div>

        <div className="mt-10 flex justify-end">
          <button
            type="button"
            className="rounded-2xl bg-gray-900 px-9 py-5 text-sm font-bold uppercase tracking-[0.18em] text-white shadow-xl shadow-gray-900/20 transition-all hover:-translate-y-0.5 hover:bg-brand-orange hover:shadow-brand-orange/25 active:translate-y-0"
          >
            Save changes
          </button>
        </div>
      </div>
    </section>
  );
}
