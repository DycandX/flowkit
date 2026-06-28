#!/usr/bin/env node
const { spawn } = require("child_process");
const { existsSync, mkdirSync, createWriteStream, createReadStream, unlinkSync, chmodSync } = require("fs");
const { get } = require("https");
const { createGunzip } = require("zlib");
const path = require("path");
const os = require("os");

const PKG = require("../package.json");
const REPO = "DycandX/flowkit";
const BIN_DIR = path.join(__dirname, "..", ".bin");
const VERSION = PKG.version;

function platformAsset() {
  const plat = os.platform();
  const arch = os.arch();
  const map = {
    "win32-x64": { file: `flowkit_${VERSION}_windows_amd64.zip`, bin: "flowkit.exe" },
    "darwin-x64": { file: `flowkit_${VERSION}_darwin_amd64.tar.gz`, bin: "flowkit" },
    "darwin-arm64": { file: `flowkit_${VERSION}_darwin_arm64.tar.gz`, bin: "flowkit" },
    "linux-x64": { file: `flowkit_${VERSION}_linux_amd64.tar.gz`, bin: "flowkit" },
    "linux-arm64": { file: `flowkit_${VERSION}_linux_arm64.tar.gz`, bin: "flowkit" },
  };
  const key = `${plat}-${arch}`;
  if (!map[key]) {
    console.error(`Unsupported platform: ${plat} ${arch}`);
    console.error(`Install via: go install github.com/${REPO}@latest`);
    process.exit(1);
  }
  return map[key];
}

function binPath() {
  return path.join(BIN_DIR, platformAsset().bin);
}

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = createWriteStream(dest);
    get(url, (res) => {
      if (res.statusCode !== 200) {
        reject(new Error(`Download failed: HTTP ${res.statusCode}`));
        return;
      }
      res.pipe(file);
      file.on("finish", () => { file.close(); resolve(); });
    }).on("error", (err) => { try { unlinkSync(dest); } catch {} reject(err); });
  });
}

function extractZip(zipPath, destDir) {
  const AdmZip = require("adm-zip");
  return new Promise((resolve, reject) => {
    try {
      const zip = new AdmZip(zipPath);
      zip.extractAllTo(destDir, true);
      resolve();
    } catch (e) { reject(e); }
  });
}

function extractTarGz(tgzPath, destDir, entryName) {
  const tar = require("tar-stream").extract();
  const out = createWriteStream(path.join(destDir, entryName));
  return new Promise((resolve, reject) => {
    tar.on("entry", (header, stream, next) => {
      stream.resume();
      if (header.name === entryName || header.name.endsWith("/" + entryName)) {
        stream.pipe(out);
        out.on("finish", next);
      } else {
        stream.on("end", next);
      }
    });
    tar.on("finish", resolve);
    tar.on("error", reject);
    createReadStream(tgzPath).pipe(createGunzip()).pipe(tar);
  });
}

async function ensureBinary() {
  const bPath = binPath();
  if (existsSync(bPath)) return bPath;
  if (!existsSync(BIN_DIR)) mkdirSync(BIN_DIR, { recursive: true });

  const asset = platformAsset();
  const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${asset.file}`;
  const tmp = path.join(BIN_DIR, `download-${Date.now()}`);

  console.error(`Downloading flowkit v${VERSION}...`);
  await download(url, tmp);

  if (asset.file.endsWith(".zip")) {
    await extractZip(tmp, BIN_DIR);
    unlinkSync(tmp);
  } else {
    await extractTarGz(tmp, BIN_DIR, asset.bin);
    unlinkSync(tmp);
  }
  try { chmodSync(bPath, 0o755); } catch {}
  return bPath;
}

async function main() {
  try {
    const bPath = await ensureBinary();
    const child = spawn(bPath, process.argv.slice(2), { stdio: "inherit" });
    child.on("exit", (code) => process.exit(code));
  } catch (err) {
    console.error("flowkit:", err.message);
    process.exit(1);
  }
}

main();
