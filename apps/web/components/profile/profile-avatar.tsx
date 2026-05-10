/* eslint-disable @next/next/no-img-element */

import { Plus } from "lucide-react";

interface ProfileAvatarProps {
  avatarUrl?: string;
}

export function ProfileAvatar({ avatarUrl }: ProfileAvatarProps) {
  const imageSource = avatarUrl?.trim() || "/images/community-4.jpg";

  return (
    <div className="relative h-28 w-28 shrink-0 rounded-[1.5rem] border-6 border-orange-100 bg-orange-50 shadow-xl shadow-orange-900/10">
      <img
        src={imageSource}
        alt="Profile photo"
        className="h-full w-full rounded-[1.15rem] object-cover"
      />
      <div
        className="absolute -bottom-2 -right-2 flex h-10 w-10 items-center justify-center rounded-xl border-4 border-white bg-brand-orange text-white shadow-lg shadow-brand-orange/25"
        aria-hidden="true"
      >
        <Plus className="h-5 w-5" />
      </div>
    </div>
  );
}
