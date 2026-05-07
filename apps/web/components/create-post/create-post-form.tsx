"use client";

import { useEffect, useMemo, useState } from "react";
import Image from "next/image";
import { ImagePlus, X } from "lucide-react";
import { CommunityShareSection } from "@/components/create-post/community-share-section";
import { Field, inputClassName } from "@/components/create-post/field";
import { FormSection } from "@/components/create-post/form-section";
import { SegmentedControl } from "@/components/create-post/segmented-control";
import { Stepper } from "@/components/create-post/stepper";
import { ToggleRow } from "@/components/create-post/toggle-row";
import { SelectControl } from "@/components/forms/select-control";

const termLengthOptions = ["4 months", "8 months", "12 months"];
const termSeasonOptions = ["Fall", "Spring", "Winter"];
const maxListingImages = 5;

export function CreatePostForm() {
  const [isSublet, setIsSublet] = useState(false);
  const [selectedTermLengths, setSelectedTermLengths] = useState<string[]>([]);
  const [selectedTermSeasons, setSelectedTermSeasons] = useState<string[]>([]);
  const [listingImages, setListingImages] = useState<File[]>([]);
  const imagePreviews = useMemo(
    () =>
      listingImages.map((image) => ({
        image,
        previewUrl: URL.createObjectURL(image),
      })),
    [listingImages],
  );

  useEffect(() => {
    return () => {
      imagePreviews.forEach(({ previewUrl }) => URL.revokeObjectURL(previewUrl));
    };
  }, [imagePreviews]);

  const toggleTermLength = (option: string) => {
    setSelectedTermLengths((current) =>
      current.includes(option)
        ? current.filter((value) => value !== option)
        : [...current, option],
    );
  };

  const toggleTermSeason = (option: string) => {
    setSelectedTermSeasons((current) =>
      current.includes(option)
        ? current.filter((value) => value !== option)
        : [...current, option],
    );
  };

  const handleImageUpload = (files: FileList | null) => {
    if (!files) {
      return;
    }

    setListingImages((current) =>
      [...current, ...Array.from(files)].slice(0, maxListingImages),
    );
  };

  const removeListingImage = (imageIndex: number) => {
    setListingImages((current) =>
      current.filter((_, index) => index !== imageIndex),
    );
  };

  return (
    <form className="space-y-10">
      <FormSection title="Essential Information">
        <div className="space-y-6">
          <Field label="Listing Title">
            <input
              className={inputClassName}
              placeholder="e.g. Spacious Studio near St. George Campus"
            />
          </Field>

          <Field label="Address">
            <input className={inputClassName} placeholder="Full address of the property" />
          </Field>

          <div className="grid gap-6 md:grid-cols-2">
            <Field label="Price per Month ($)">
              <input className={inputClassName} type="number" defaultValue={1000} />
            </Field>

            <Field label="Property Type">
              <SelectControl
                label="Property Type"
                options={["Apartment", "House", "Studio", "Shared Room", "Dorm"]}
                variant="create"
              />
            </Field>
          </div>
        </div>
      </FormSection>

      <FormSection title="Room & Unit Specs">
        <div className="space-y-8">
          <div className="grid gap-6 md:grid-cols-3">
            <Stepper label="Total Rooms" initialValue={1} min={1} />
            <Stepper label="Rooms Occupied" initialValue={0} />
            <Stepper label="Bathrooms" initialValue={1} min={1} />
          </div>

          <SegmentedControl
            label="Gender"
            options={["Female Only", "Male Only", "Coed"]}
            defaultValue="Coed"
            size="sm"
          />

          <ToggleRow enabled={isSublet} onChange={setIsSublet} />
        </div>
      </FormSection>

      {isSublet && (
        <FormSection title="Term Details">
          <div className="space-y-8">
            <MultiSelectPills
              label="Term Season"
              options={termSeasonOptions}
              selectedOptions={selectedTermSeasons}
              onToggle={toggleTermSeason}
            />

            <MultiSelectPills
              label="Term Length"
              options={termLengthOptions}
              selectedOptions={selectedTermLengths}
              onToggle={toggleTermLength}
            />
          </div>
        </FormSection>
      )}

      <div className="space-y-7">
        <Field label="Quick Description (for card)">
          <textarea
            className={`${inputClassName} min-h-28 resize-y`}
            placeholder="Short summary of the housing..."
          />
        </Field>

        <Field label="Detailed Content">
          <textarea
            className={`${inputClassName} min-h-56 resize-y`}
            placeholder="Describe your property, rules, roommates, community vibes, etc..."
          />
        </Field>

        <Field label="Photos">
          <div className="space-y-4">
            <label className="flex min-h-36 cursor-pointer flex-col items-center justify-center rounded-2xl border border-dashed border-orange-200 bg-orange-50/30 px-6 py-8 text-center transition-all hover:border-brand-orange hover:bg-orange-50/60">
              <ImagePlus className="mb-3 h-8 w-8 text-brand-orange" />
              <span className="text-base font-bold text-gray-700">
                Upload property photos
              </span>
              <span className="mt-1 text-sm font-medium text-gray-400">
                Add up to {maxListingImages} images
              </span>
              <input
                type="file"
                accept="image/*"
                multiple
                className="sr-only"
                onChange={(event) => handleImageUpload(event.target.files)}
              />
            </label>

            {imagePreviews.length > 0 && (
              <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-5">
                {imagePreviews.map(({ image, previewUrl }, index) => (
                  <div
                    key={`${image.name}-${index}`}
                    className="relative rounded-2xl border border-gray-100 bg-white p-3 shadow-sm"
                  >
                    <button
                      type="button"
                      onClick={() => removeListingImage(index)}
                      className="absolute right-1.5 top-1.5 rounded-full bg-white p-1 text-gray-400 shadow-md transition-colors hover:text-brand-orange"
                      aria-label={`Remove ${image.name}`}
                    >
                      <X size={16} />
                    </button>
                    <Image
                      src={previewUrl}
                      alt={image.name}
                      width={160}
                      height={160}
                      unoptimized
                      className="mb-2 aspect-square w-full rounded-xl object-cover"
                    />
                    <p className="truncate text-xs font-bold text-gray-600">
                      {image.name}
                    </p>
                  </div>
                ))}
              </div>
            )}
          </div>
        </Field>
      </div>

      <CommunityShareSection />

      <div className="mb-10 flex justify-end">
        <button className="rounded-2xl bg-brand-orange px-10 py-4 text-base font-black uppercase tracking-[0.18em] text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99]">
          Submit
        </button>
      </div>
    </form>
  );
}

interface MultiSelectPillsProps {
  label: string;
  options: string[];
  selectedOptions: string[];
  onToggle: (option: string) => void;
}

function MultiSelectPills({
  label,
  options,
  selectedOptions,
  onToggle,
}: MultiSelectPillsProps) {
  return (
    <div>
      <p className="mb-2 text-sm font-bold uppercase tracking-[0.16em] text-gray-400">
        {label}
      </p>
      <div
        className="grid rounded-2xl border border-gray-100 bg-white p-1 shadow-sm"
        style={{ gridTemplateColumns: `repeat(${options.length}, minmax(0, 1fr))` }}
      >
        {options.map((option) => {
          const isSelected = selectedOptions.includes(option);

          return (
            <button
              key={option}
              type="button"
              onClick={() => onToggle(option)}
              aria-pressed={isSelected}
              className={`rounded-xl px-4 py-3 text-sm font-bold uppercase transition-all ${
                isSelected
                  ? "bg-brand-orange text-white shadow-lg shadow-brand-orange/20"
                  : "text-gray-400 hover:text-brand-orange"
              }`}
            >
              {option}
            </button>
          );
        })}
      </div>
    </div>
  );
}
