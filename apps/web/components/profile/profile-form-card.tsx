import { ProfileAvatar } from "@/components/profile/profile-avatar";
import { ProfileContactCard } from "@/components/profile/profile-contact-card";
import {
  FieldLabel,
  GenderSelect,
  ProfileField,
} from "@/components/profile/profile-field";

const profileFields = [
  { label: "First Name", value: "Alex" },
  { label: "Last Name", value: "Rivera" },
  { label: "Age", value: "21", type: "number" },
  { label: "Institution", value: "University of Toronto" },
  { label: "Major", value: "Architecture & Design" },
];

export function ProfileFormCard() {
  return (
    <section className="rounded-[2rem] border border-orange-100/70 bg-white px-6 py-8 shadow-2xl shadow-orange-900/8 sm:px-8 lg:px-12 lg:py-12">
      <div className="flex flex-col gap-5 sm:flex-row sm:items-end">
        <ProfileAvatar />

        <div className="pb-2">
          <h1 className="text-4xl font-bold tracking-tight text-gray-900">
            Alex Rivera
          </h1>
          <p className="mt-2 text-lg font-semibold italic text-gray-400">
            Student at University of Toronto
          </p>
        </div>
      </div>

      <div className="mt-10 grid gap-6 md:grid-cols-2">
        <ProfileField label={profileFields[0].label} value={profileFields[0].value} />
        <ProfileField label={profileFields[1].label} value={profileFields[1].value} />
        <GenderSelect />
        <ProfileField
          label={profileFields[2].label}
          value={profileFields[2].value}
          type={profileFields[2].type}
        />
        <ProfileField label={profileFields[3].label} value={profileFields[3].value} />
        <ProfileField label={profileFields[4].label} value={profileFields[4].value} />
      </div>

      <div className="mt-6">
        <label className="block">
          <FieldLabel>Bio</FieldLabel>
          <textarea
            defaultValue="Looking for clean, transit-friendly housing near campus. I like quiet study spaces, shared dinners, and communities that keep things organized."
            rows={5}
            className="w-full resize-none rounded-3xl border border-gray-100 bg-gray-50 px-6 py-5 text-base font-semibold leading-relaxed text-gray-700 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10"
          />
        </label>
      </div>

      <div className="mt-6">
        <ProfileContactCard />
      </div>

      <div className="mt-10 flex justify-end">
        <button
          type="button"
          className="rounded-2xl bg-gray-900 px-9 py-5 text-sm font-bold uppercase tracking-[0.18em] text-white shadow-xl shadow-gray-900/20 transition-all hover:-translate-y-0.5 hover:bg-brand-orange hover:shadow-brand-orange/25 active:translate-y-0"
        >
          Save changes
        </button>
      </div>
    </section>
  );
}
