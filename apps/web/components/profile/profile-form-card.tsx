import type { ChangeEvent } from "react";
import { ProfileAvatar } from "@/components/profile/profile-avatar";
import { ProfileContactCard } from "@/components/profile/profile-contact-card";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";
import {
  FieldLabel,
  GenderSelect,
  ProfileField,
} from "@/components/profile/profile-field";

type InstitutionOption = {
  id: string;
  name: string;
};

export type ProfileFormState = {
  username: string;
  email: string;
  avatar: string;
  bio: string;
  gender: string;
  age: string;
  institutionId: string;
  major: string;
};

interface ProfileFormCardProps {
  form: ProfileFormState;
  institutions: InstitutionOption[];
  statusMessage?: string;
  errorMessage?: string;
  isSaving?: boolean;
  onFieldChange: (field: keyof ProfileFormState, value: string) => void;
  onSave: () => void;
}

export function ProfileFormCard({
  form,
  institutions,
  statusMessage,
  errorMessage,
  isSaving = false,
  onFieldChange,
  onSave,
}: ProfileFormCardProps) {
  const selectedInstitution = institutions.find(
    (institution) => institution.id === form.institutionId,
  )?.name;

  const handleBioChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    onFieldChange("bio", event.target.value);
  };

  return (
    <section className={`${profilePanelClassName} overflow-y-auto`}>
      <div className="flex flex-col gap-4 sm:flex-row sm:items-end">
        <ProfileAvatar avatarUrl={form.avatar} />

        <div className="pb-1">
          <h1 className="text-3xl font-bold tracking-tight text-gray-900">
            {form.username || "Your profile"}
          </h1>
          <p className="mt-1.5 text-base font-semibold italic text-gray-500">
            {selectedInstitution
              ? `Student at ${selectedInstitution}`
              : "Add your institution to personalize your profile"}
          </p>
        </div>
      </div>

      <div className="mt-6 grid gap-x-5 gap-y-4 md:grid-cols-2">
        <ProfileField
          label="Name"
          value={form.username}
          onChange={(value) => onFieldChange("username", value)}
        />
        <GenderSelect
          value={form.gender}
          onChange={(value) => onFieldChange("gender", value)}
        />
        <ProfileField
          label="Age (optional)"
          value={form.age}
          onChange={(value) => onFieldChange("age", value)}
          type="number"
          min={18}
          max={100}
        />
        <label className="block">
          <FieldLabel>Institution</FieldLabel>
          <select
            value={form.institutionId}
            onChange={(event) => onFieldChange("institutionId", event.target.value)}
            className="h-12 w-full appearance-none rounded-2xl border border-gray-100 bg-gray-50 px-5 text-sm font-bold text-gray-800 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10"
          >
            <option value="">Select an institution</option>
            {institutions.map((institution) => (
              <option key={institution.id} value={institution.id}>
                {institution.name}
              </option>
            ))}
          </select>
        </label>
        <ProfileField
          label="Major (optional)"
          value={form.major}
          onChange={(value) => onFieldChange("major", value)}
        />
        <ProfileContactCard
          email={form.email}
          onChange={(value) => onFieldChange("email", value)}
        />
        <ProfileField
          label="Avatar URL (optional)"
          value={form.avatar}
          onChange={(value) => onFieldChange("avatar", value)}
          placeholder="https://example.com/avatar.jpg"
        />
      </div>

      <div className="mt-4">
        <label className="block">
          <FieldLabel>Bio</FieldLabel>
          <textarea
            value={form.bio}
            onChange={handleBioChange}
            rows={5}
            className="w-full resize-none rounded-2xl border border-gray-100 bg-gray-50 px-5 py-4 text-sm font-semibold leading-6 text-gray-700 shadow-inner shadow-gray-900/5 outline-none transition-all focus:border-brand-orange focus:bg-white focus:ring-4 focus:ring-brand-orange/10"
          />
        </label>
      </div>

      {errorMessage ? (
        <p className="mt-4 rounded-2xl border border-red-200 bg-red-50 px-4 py-3 text-sm font-semibold text-red-700">
          {errorMessage}
        </p>
      ) : null}
      {statusMessage ? (
        <p className="mt-4 rounded-2xl border border-orange-200 bg-orange-50 px-4 py-3 text-sm font-semibold text-orange-700">
          {statusMessage}
        </p>
      ) : null}

      <div className="mt-6 flex justify-end">
        <button
          type="button"
          onClick={onSave}
          disabled={isSaving}
          className="rounded-2xl bg-brand-orange px-7 py-4 text-xs font-bold uppercase tracking-[0.16em] text-white shadow-xl shadow-brand-orange/25 transition-all hover:-translate-y-0.5 hover:bg-orange-600 active:translate-y-0"
        >
          {isSaving ? "Saving..." : "Save changes"}
        </button>
      </div>
    </section>
  );
}
