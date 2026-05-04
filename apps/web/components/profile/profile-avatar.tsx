import Image from "next/image";
import { Plus } from "lucide-react";

export function ProfileAvatar() {
  return (
    <div className="relative h-28 w-28 shrink-0 rounded-[1.5rem] border-6 border-orange-50 bg-orange-50 shadow-xl shadow-orange-900/10">
      <Image
        src="/images/community-4.jpg"
        alt="Alex Rivera profile photo"
        fill
        sizes="112px"
        className="rounded-[1.15rem] object-cover"
      />
      <button
        type="button"
        className="absolute -bottom-2 -right-2 flex h-10 w-10 items-center justify-center rounded-xl border-4 border-white bg-gray-900 text-white shadow-lg transition-transform hover:scale-105"
        aria-label="Upload profile photo"
      >
        <Plus className="h-5 w-5" />
      </button>
    </div>
  );
}
