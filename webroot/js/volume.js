// Volume Control Functions
function setVolumeStatus(status) {
    const el = document.getElementById('volumeStatus');
    el.innerHTML = `<p>${status}</p>`;
}

async function getVolume() {
    try {
        const res = await fetch('/api/mods/VolumeControl/get');
        if (!res.ok) {
            const e = await res.json();
            console.log(`Volume get failed: ${e.message}`);
            document.getElementById('volumeSelect').value = 'MEDIUM'; // Default fallback
        } else {
            const volume = await res.text();
            document.getElementById('volumeSelect').value = volume;
        }
    } catch (e) {
        console.log(`Volume network error: ${e.message}`);
        document.getElementById('volumeSelect').value = 'MEDIUM'; // Default fallback
    }
}

async function setVolume() {
    const level = document.getElementById('volumeSelect').value;
    setVolumeStatus('Setting volume...');
    
    try {
        const res = await fetch(`/api/mods/VolumeControl/set?level=${level}`);
        if (!res.ok) {
            const e = await res.json();
            setVolumeStatus(`Failed to set volume: ${e.message}`);
        } else {
            setVolumeStatus(`Volume set to: ${level}`);
            // Clear status after 3 seconds
            setTimeout(() => {
                document.getElementById('volumeStatus').innerHTML = '';
            }, 3000);
        }
    } catch (e) {
        setVolumeStatus(`Network error: ${e.message}`);
    }
}